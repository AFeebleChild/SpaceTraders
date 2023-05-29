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
	content, err := JsonReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot load waypoints file: %v", err)
	}
	waypoints := &Waypoints{}
	err = json.Unmarshal(content, waypoints)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal waypoints file: %v", err)
	}

	return waypoints, nil
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
