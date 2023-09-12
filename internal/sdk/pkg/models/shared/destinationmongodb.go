// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationMongodbAuthorizationTypeLoginPasswordAuthorization string

const (
	DestinationMongodbAuthorizationTypeLoginPasswordAuthorizationLoginPassword DestinationMongodbAuthorizationTypeLoginPasswordAuthorization = "login/password"
)

func (e DestinationMongodbAuthorizationTypeLoginPasswordAuthorization) ToPointer() *DestinationMongodbAuthorizationTypeLoginPasswordAuthorization {
	return &e
}

func (e *DestinationMongodbAuthorizationTypeLoginPasswordAuthorization) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "login/password":
		*e = DestinationMongodbAuthorizationTypeLoginPasswordAuthorization(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbAuthorizationTypeLoginPasswordAuthorization: %v", v)
	}
}

// DestinationMongodbAuthorizationTypeLoginPassword - Login/Password.
type DestinationMongodbAuthorizationTypeLoginPassword struct {
	Authorization DestinationMongodbAuthorizationTypeLoginPasswordAuthorization `json:"authorization"`
	// Password associated with the username.
	Password string `json:"password"`
	// Username to use to access the database.
	Username string `json:"username"`
}

type DestinationMongodbAuthorizationTypeNoneAuthorization string

const (
	DestinationMongodbAuthorizationTypeNoneAuthorizationNone DestinationMongodbAuthorizationTypeNoneAuthorization = "none"
)

func (e DestinationMongodbAuthorizationTypeNoneAuthorization) ToPointer() *DestinationMongodbAuthorizationTypeNoneAuthorization {
	return &e
}

func (e *DestinationMongodbAuthorizationTypeNoneAuthorization) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		*e = DestinationMongodbAuthorizationTypeNoneAuthorization(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbAuthorizationTypeNoneAuthorization: %v", v)
	}
}

// DestinationMongodbAuthorizationTypeNone - None.
type DestinationMongodbAuthorizationTypeNone struct {
	Authorization DestinationMongodbAuthorizationTypeNoneAuthorization `json:"authorization"`
}

type DestinationMongodbAuthorizationTypeType string

const (
	DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeNone          DestinationMongodbAuthorizationTypeType = "destination-mongodb_Authorization type_None"
	DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeLoginPassword DestinationMongodbAuthorizationTypeType = "destination-mongodb_Authorization type_Login/Password"
)

type DestinationMongodbAuthorizationType struct {
	DestinationMongodbAuthorizationTypeNone          *DestinationMongodbAuthorizationTypeNone
	DestinationMongodbAuthorizationTypeLoginPassword *DestinationMongodbAuthorizationTypeLoginPassword

	Type DestinationMongodbAuthorizationTypeType
}

func CreateDestinationMongodbAuthorizationTypeDestinationMongodbAuthorizationTypeNone(destinationMongodbAuthorizationTypeNone DestinationMongodbAuthorizationTypeNone) DestinationMongodbAuthorizationType {
	typ := DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeNone

	return DestinationMongodbAuthorizationType{
		DestinationMongodbAuthorizationTypeNone: &destinationMongodbAuthorizationTypeNone,
		Type:                                    typ,
	}
}

func CreateDestinationMongodbAuthorizationTypeDestinationMongodbAuthorizationTypeLoginPassword(destinationMongodbAuthorizationTypeLoginPassword DestinationMongodbAuthorizationTypeLoginPassword) DestinationMongodbAuthorizationType {
	typ := DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeLoginPassword

	return DestinationMongodbAuthorizationType{
		DestinationMongodbAuthorizationTypeLoginPassword: &destinationMongodbAuthorizationTypeLoginPassword,
		Type: typ,
	}
}

func (u *DestinationMongodbAuthorizationType) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMongodbAuthorizationTypeNone := new(DestinationMongodbAuthorizationTypeNone)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbAuthorizationTypeNone); err == nil {
		u.DestinationMongodbAuthorizationTypeNone = destinationMongodbAuthorizationTypeNone
		u.Type = DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeNone
		return nil
	}

	destinationMongodbAuthorizationTypeLoginPassword := new(DestinationMongodbAuthorizationTypeLoginPassword)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbAuthorizationTypeLoginPassword); err == nil {
		u.DestinationMongodbAuthorizationTypeLoginPassword = destinationMongodbAuthorizationTypeLoginPassword
		u.Type = DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeLoginPassword
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMongodbAuthorizationType) MarshalJSON() ([]byte, error) {
	if u.DestinationMongodbAuthorizationTypeNone != nil {
		return json.Marshal(u.DestinationMongodbAuthorizationTypeNone)
	}

	if u.DestinationMongodbAuthorizationTypeLoginPassword != nil {
		return json.Marshal(u.DestinationMongodbAuthorizationTypeLoginPassword)
	}

	return nil, nil
}

type DestinationMongodbMongodb string

const (
	DestinationMongodbMongodbMongodb DestinationMongodbMongodb = "mongodb"
)

func (e DestinationMongodbMongodb) ToPointer() *DestinationMongodbMongodb {
	return &e
}

func (e *DestinationMongodbMongodb) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mongodb":
		*e = DestinationMongodbMongodb(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbMongodb: %v", v)
	}
}

type DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstance string

const (
	DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstanceAtlas DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstance = "atlas"
)

func (e DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstance) ToPointer() *DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstance {
	return &e
}

func (e *DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "atlas":
		*e = DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstance: %v", v)
	}
}

// DestinationMongodbMongoDBInstanceTypeMongoDBAtlas - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type DestinationMongodbMongoDBInstanceTypeMongoDBAtlas struct {
	// URL of a cluster to connect to.
	ClusterURL string                                                    `json:"cluster_url"`
	Instance   DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstance `json:"instance"`
}

type DestinationMongodbMongoDbInstanceTypeReplicaSetInstance string

const (
	DestinationMongodbMongoDbInstanceTypeReplicaSetInstanceReplica DestinationMongodbMongoDbInstanceTypeReplicaSetInstance = "replica"
)

func (e DestinationMongodbMongoDbInstanceTypeReplicaSetInstance) ToPointer() *DestinationMongodbMongoDbInstanceTypeReplicaSetInstance {
	return &e
}

func (e *DestinationMongodbMongoDbInstanceTypeReplicaSetInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "replica":
		*e = DestinationMongodbMongoDbInstanceTypeReplicaSetInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbMongoDbInstanceTypeReplicaSetInstance: %v", v)
	}
}

// DestinationMongodbMongoDbInstanceTypeReplicaSet - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type DestinationMongodbMongoDbInstanceTypeReplicaSet struct {
	Instance DestinationMongodbMongoDbInstanceTypeReplicaSetInstance `json:"instance"`
	// A replica set name.
	ReplicaSet *string `json:"replica_set,omitempty"`
	// The members of a replica set. Please specify `host`:`port` of each member seperated by comma.
	ServerAddresses string `json:"server_addresses"`
}

type DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstance string

const (
	DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceStandalone DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstance = "standalone"
)

func (e DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstance) ToPointer() *DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstance {
	return &e
}

func (e *DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standalone":
		*e = DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstance: %v", v)
	}
}

// DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance struct {
	// The Host of a Mongo database to be replicated.
	Host     string                                                                 `json:"host"`
	Instance DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstance `json:"instance"`
	// The Port of a Mongo database to be replicated.
	Port int64 `json:"port"`
}

type DestinationMongodbMongoDbInstanceTypeType string

const (
	DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance DestinationMongodbMongoDbInstanceTypeType = "destination-mongodb_MongoDb Instance Type_Standalone MongoDb Instance"
	DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeReplicaSet                DestinationMongodbMongoDbInstanceTypeType = "destination-mongodb_MongoDb Instance Type_Replica Set"
	DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDBInstanceTypeMongoDBAtlas              DestinationMongodbMongoDbInstanceTypeType = "destination-mongodb_MongoDb Instance Type_MongoDB Atlas"
)

type DestinationMongodbMongoDbInstanceType struct {
	DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance *DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance
	DestinationMongodbMongoDbInstanceTypeReplicaSet                *DestinationMongodbMongoDbInstanceTypeReplicaSet
	DestinationMongodbMongoDBInstanceTypeMongoDBAtlas              *DestinationMongodbMongoDBInstanceTypeMongoDBAtlas

	Type DestinationMongodbMongoDbInstanceTypeType
}

func CreateDestinationMongodbMongoDbInstanceTypeDestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance(destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance) DestinationMongodbMongoDbInstanceType {
	typ := DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance

	return DestinationMongodbMongoDbInstanceType{
		DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance: &destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance,
		Type: typ,
	}
}

func CreateDestinationMongodbMongoDbInstanceTypeDestinationMongodbMongoDbInstanceTypeReplicaSet(destinationMongodbMongoDbInstanceTypeReplicaSet DestinationMongodbMongoDbInstanceTypeReplicaSet) DestinationMongodbMongoDbInstanceType {
	typ := DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeReplicaSet

	return DestinationMongodbMongoDbInstanceType{
		DestinationMongodbMongoDbInstanceTypeReplicaSet: &destinationMongodbMongoDbInstanceTypeReplicaSet,
		Type: typ,
	}
}

func CreateDestinationMongodbMongoDbInstanceTypeDestinationMongodbMongoDBInstanceTypeMongoDBAtlas(destinationMongodbMongoDBInstanceTypeMongoDBAtlas DestinationMongodbMongoDBInstanceTypeMongoDBAtlas) DestinationMongodbMongoDbInstanceType {
	typ := DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDBInstanceTypeMongoDBAtlas

	return DestinationMongodbMongoDbInstanceType{
		DestinationMongodbMongoDBInstanceTypeMongoDBAtlas: &destinationMongodbMongoDBInstanceTypeMongoDBAtlas,
		Type: typ,
	}
}

func (u *DestinationMongodbMongoDbInstanceType) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMongodbMongoDBInstanceTypeMongoDBAtlas := new(DestinationMongodbMongoDBInstanceTypeMongoDBAtlas)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbMongoDBInstanceTypeMongoDBAtlas); err == nil {
		u.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas = destinationMongodbMongoDBInstanceTypeMongoDBAtlas
		u.Type = DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDBInstanceTypeMongoDBAtlas
		return nil
	}

	destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance := new(DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance); err == nil {
		u.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance = destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance
		u.Type = DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance
		return nil
	}

	destinationMongodbMongoDbInstanceTypeReplicaSet := new(DestinationMongodbMongoDbInstanceTypeReplicaSet)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbMongoDbInstanceTypeReplicaSet); err == nil {
		u.DestinationMongodbMongoDbInstanceTypeReplicaSet = destinationMongodbMongoDbInstanceTypeReplicaSet
		u.Type = DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeReplicaSet
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMongodbMongoDbInstanceType) MarshalJSON() ([]byte, error) {
	if u.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas != nil {
		return json.Marshal(u.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas)
	}

	if u.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
		return json.Marshal(u.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance)
	}

	if u.DestinationMongodbMongoDbInstanceTypeReplicaSet != nil {
		return json.Marshal(u.DestinationMongodbMongoDbInstanceTypeReplicaSet)
	}

	return nil, nil
}

// DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationMongodbSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationMongodbSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationMongodbSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethod `json:"tunnel_method"`
}

type DestinationMongodbSSHTunnelMethodType string

const (
	DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodNoTunnel               DestinationMongodbSSHTunnelMethodType = "destination-mongodb_SSH Tunnel Method_No Tunnel"
	DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodSSHKeyAuthentication   DestinationMongodbSSHTunnelMethodType = "destination-mongodb_SSH Tunnel Method_SSH Key Authentication"
	DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodPasswordAuthentication DestinationMongodbSSHTunnelMethodType = "destination-mongodb_SSH Tunnel Method_Password Authentication"
)

type DestinationMongodbSSHTunnelMethod struct {
	DestinationMongodbSSHTunnelMethodNoTunnel               *DestinationMongodbSSHTunnelMethodNoTunnel
	DestinationMongodbSSHTunnelMethodSSHKeyAuthentication   *DestinationMongodbSSHTunnelMethodSSHKeyAuthentication
	DestinationMongodbSSHTunnelMethodPasswordAuthentication *DestinationMongodbSSHTunnelMethodPasswordAuthentication

	Type DestinationMongodbSSHTunnelMethodType
}

func CreateDestinationMongodbSSHTunnelMethodDestinationMongodbSSHTunnelMethodNoTunnel(destinationMongodbSSHTunnelMethodNoTunnel DestinationMongodbSSHTunnelMethodNoTunnel) DestinationMongodbSSHTunnelMethod {
	typ := DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodNoTunnel

	return DestinationMongodbSSHTunnelMethod{
		DestinationMongodbSSHTunnelMethodNoTunnel: &destinationMongodbSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationMongodbSSHTunnelMethodDestinationMongodbSSHTunnelMethodSSHKeyAuthentication(destinationMongodbSSHTunnelMethodSSHKeyAuthentication DestinationMongodbSSHTunnelMethodSSHKeyAuthentication) DestinationMongodbSSHTunnelMethod {
	typ := DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodSSHKeyAuthentication

	return DestinationMongodbSSHTunnelMethod{
		DestinationMongodbSSHTunnelMethodSSHKeyAuthentication: &destinationMongodbSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationMongodbSSHTunnelMethodDestinationMongodbSSHTunnelMethodPasswordAuthentication(destinationMongodbSSHTunnelMethodPasswordAuthentication DestinationMongodbSSHTunnelMethodPasswordAuthentication) DestinationMongodbSSHTunnelMethod {
	typ := DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodPasswordAuthentication

	return DestinationMongodbSSHTunnelMethod{
		DestinationMongodbSSHTunnelMethodPasswordAuthentication: &destinationMongodbSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationMongodbSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMongodbSSHTunnelMethodNoTunnel := new(DestinationMongodbSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationMongodbSSHTunnelMethodNoTunnel = destinationMongodbSSHTunnelMethodNoTunnel
		u.Type = DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodNoTunnel
		return nil
	}

	destinationMongodbSSHTunnelMethodSSHKeyAuthentication := new(DestinationMongodbSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication = destinationMongodbSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationMongodbSSHTunnelMethodPasswordAuthentication := new(DestinationMongodbSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationMongodbSSHTunnelMethodPasswordAuthentication = destinationMongodbSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMongodbSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMongodbSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationMongodbSSHTunnelMethodNoTunnel)
	}

	if u.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationMongodbSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationMongodbSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type DestinationMongodb struct {
	// Authorization type.
	AuthType DestinationMongodbAuthorizationType `json:"auth_type"`
	// Name of the database.
	Database        string                    `json:"database"`
	DestinationType DestinationMongodbMongodb `json:"destinationType"`
	// MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
	InstanceType *DestinationMongodbMongoDbInstanceType `json:"instance_type,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMongodbSSHTunnelMethod `json:"tunnel_method,omitempty"`
}
