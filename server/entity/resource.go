package entity

type Resource struct {
	DBModel
	ClusterName string
	NodeIP      string
}

type Resources struct {
	ClusterName  string
	NodeIP       string
	DaemonUpdate string
	FalcoUpdate  string
}
