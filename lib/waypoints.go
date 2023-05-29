package lib

import (
	"context"
	"encoding/json"
	"fmt"
)

type (
	Waypoints []Waypoint
)

func (w Waypoints) Save(system string) error {
	path := "systems/" + system + "/waypoints.json"
	fmt.Println("Saving waypoints: ", path)

	return JsonFilePrettyPrint(path, w)
}

func LoadWaypoints(system string) (*Waypoints, error) {
	path := "systems/" + system + "/waypoints.json"
	waypoints := &Waypoints{}
	err := JsonReadFile(path, waypoints)
	return waypoints, err
}

func GetWaypoints(c *Client, system string) (*Waypoints, error) {
	params := &GetSystemWaypointsParams{}

	body, err := HandleResp(c.GetSystemWaypoints(context.Background(), system, params))
	if err != nil {
		return nil, err
	}

	waypoints := &Waypoints{}
	err = json.Unmarshal(body, waypoints)
	if err != nil {
		return nil, err
	}

	return waypoints, nil
}
