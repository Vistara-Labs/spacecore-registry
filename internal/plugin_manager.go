package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"spacecore_registry/pb"
	"strings"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	// "github.com/libp2p/go-libp2p/p2p/discovery/routing"
)

type pluginRegistryServer struct {
	pb.UnimplementedPluginRegistryServer
	ipfsClient  *rpc.HttpApi
	redisClient *redis.Client
	// dht *kbucket.RoutingTable
	pinataAPIKey    string
	pinataAPISecret string
}

// Define a struct for the Pinata pin request body
type pinataPinRequest struct {
	HashToPin string `json:"hashToPin"`
}

func NewPluginRegistryServer(ipfsClient *rpc.HttpApi, redisClient *redis.Client, pinataAPIKey string, pinataAPISecret string) *pluginRegistryServer {
	return &pluginRegistryServer{
		ipfsClient:      ipfsClient,
		redisClient:     redisClient,
		pinataAPIKey:    pinataAPIKey,
		pinataAPISecret: pinataAPISecret,
	}
}

// Example: vimana register ipfst /Users/mayurchougule/development/spacecore-plugins/ipfs-plugin/bin/ipfspd
func (s *pluginRegistryServer) RegisterPlugin(ctx context.Context, req *pb.RegisterPluginRequest) (*pb.RegisterPluginResponse, error) {
	log.Printf("req plugin: %v\n", req.Plugin)
	// cid, err := s.ipfsClient.Unixfs().Add(ctx, fileNode)

	cid, err := s.ipfsClient.Unixfs().Add(ctx, getUnixfsNode(req.Plugin))
	cidPath, err := path.NewPath(req.Plugin)
	log.Printf("cid: %v\n", cid)
	log.Printf("cidPath: %v\n", cidPath)

	// Assuming `cid` is the CID of the content you want to pin

	err = s.ipfsClient.Pin().Add(ctx, cid)
	if err != nil {
		return nil, fmt.Errorf("failed to pin content locally: %w", err)
	}
	log.Printf("Successfully pinned CID: %s locally\n", cid)

	gatewayUrl, _ := pinToPinata(cid.String())

	plugin := &pb.Plugin{
		Name:    req.Name,
		Version: req.Version,
		Cid:     cid.String(),
	}

	// Use the CID as the key in the DHT
	// // Provide the CID in the DHT - cannot use putvalue because its intended use is a bit different
	// // https://github.com/libp2p/go-libp2p-kad-dht/issues/381
	// if err := s.dht.Provide(ctx, cid.RootCid(), true); err != nil {
	// 	return nil, fmt.Errorf("failed to provide CID: %w", err)
	// }

	key := fmt.Sprintf("plugin:%s:%s", req.Name, req.Version)
	value, err := json.Marshal(plugin)
	if err != nil {
		return nil, err
	}

	if err := s.redisClient.Set(ctx, key, value, 0).Err(); err != nil {
		return nil, err
	}

	log.Printf("Registered Plugin: %s with CID: %s", req.Name, gatewayUrl)
	return &pb.RegisterPluginResponse{
		Message: "Plugin registered successfully",
		Cid:     gatewayUrl,
	}, nil

}

func (s *pluginRegistryServer) DiscoverPlugins(ctx context.Context, req *pb.DiscoverPluginsRequest) (*pb.DiscoverPluginsResponse, error) {
	plugins := []*pb.Plugin{}

	// if req.Name is empty, then list top 10 plugins
	if *req.Name == "" {
		searchKey := "plugin:*"
		// Use the Keys method to find all keys matching the pattern
		keys, err := s.redisClient.Keys(ctx, searchKey).Result()
		if err != nil {
			return nil, err
		}
		// If no specific search, optionally limit to top 10 keys
		if *req.Name == "" && len(keys) > 10 {
			keys = keys[:10]
		}
		for _, key := range keys {
			value, err := s.redisClient.Get(ctx, key).Result()
			if err != nil {
				// Handle error (e.g., key might have expired between fetching keys and values)
				fmt.Printf("Error fetching value for key %s: %v\n", key, err)
				continue
			}

			var plugin pb.Plugin
			if err := json.Unmarshal([]byte(value), &plugin); err != nil {
				log.Printf("failed to unmarshal plugin: %v", err)
				return nil, err
			}
			plugins = append(plugins, &plugin)
		}

		// can return the list of keys since we are not searching for a specific plugin
		return &pb.DiscoverPluginsResponse{
			Plugins: plugins,
		}, nil
	}

	key := fmt.Sprintf("plugin:%s:%s", *req.Name, "1.0")

	value, err := s.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	log.Printf("value: %v\n", value)
	var plugin pb.Plugin
	if err := json.Unmarshal([]byte(value), &plugin); err != nil {
		log.Printf("failed to unmarshal plugin: %v", err)
		return nil, err
	}
	// providers := s.dht.FindProvidersAsync(ctx, cid, 5)
	// responses, err := s.dht.SearchValue(ctx, *req.Name)
	// value, err := s.dht.GetValue(ctx, *req.Name)

	if err != nil {
		return nil, fmt.Errorf("failed to get value: %w", err)
	}
	return &pb.DiscoverPluginsResponse{
		Plugins: []*pb.Plugin{&plugin},
	}, nil
}

func getUnixfsNode(path string) files.Node {
	st, err := os.Stat(path)
	if err != nil {
		return nil
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil
	}
	return f
}

func (s *pluginRegistryServer) DownloadPlugin(ctx context.Context, req *pb.DownloadPluginRequest) (*pb.DownloadPluginResponse, error) {
	pluginPath, err := path.NewPath(req.Cid)
	if err != nil {
		return nil, err
	}
	res, err := s.ipfsClient.Unixfs().Get(ctx, pluginPath)
	if err != nil {
		return nil, err
	}

	fileReader, ok := res.(files.File)
	if !ok {
		return nil, fmt.Errorf("failed to convert fileNode type %T to files.File", res)
	}

	data, err := io.ReadAll(fileReader)
	if err != nil {
		return nil, err
	}

	return &pb.DownloadPluginResponse{
		Content: data,
	}, nil
}

func (s *pluginRegistryServer) GetPlugin(ctx context.Context, req *pb.GetPluginRequest) (*pb.GetPluginResponse, error) {
	key := fmt.Sprintf("plugin:%s:%s", req.Name, req.Version)
	value, err := s.redisClient.Get(ctx, key).Result()
	if err != nil {
		return &pb.GetPluginResponse{}, err
	}

	var plugin pb.Plugin
	if err := json.Unmarshal([]byte(value), &plugin); err != nil {
		return &pb.GetPluginResponse{}, err
	}

	return &pb.GetPluginResponse{
		Plugin: &plugin,
	}, nil
}

func pinToPinata(hashToPin string) (string, error) {
	err := godotenv.Load()
	// filter out /ipfs/ from the hash

	cidHash := strings.Replace(hashToPin, "/ipfs/", "", 1)

	url := "https://api.pinata.cloud/pinning/pinByHash"
	body := map[string]interface{}{
		"hashToPin": cidHash,
		"pinataOptions": map[string]interface{}{
			"hostNodes": []string{
				"/ip4/73.227.56.113/tcp/18667/p2p/12D3KooWJU4TpmDokPbTftD1PsSfuyx8jvysHpS5o9fL224xYKZM",
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "Error parsing body " + hashToPin, err
	}
	payload := strings.NewReader(string(jsonBody))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "Error pinning file to Pinata " + hashToPin, err
	}

	gatewayUrl := "https://gateway.pinata.cloud/ipfs/"
	bearerToken := os.Getenv("JWT")
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")
	// fmt.Printf("bearer : %v\n", bearerToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("failed to pin: %s", hashToPin), err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("failed to pin: %s", resp.Status), nil
	}

	return fmt.Sprintf("%s%s", gatewayUrl, hashToPin), nil
}
