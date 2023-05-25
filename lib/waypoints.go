package lib

import (
	"context"
	"encoding/json"
)

type (
	Waypoints []Waypoint
)

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
