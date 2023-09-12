// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationMongodbUpdateAuthorizationTypeLoginPasswordAuthorization string

const (
	DestinationMongodbUpdateAuthorizationTypeLoginPasswordAuthorizationLoginPassword DestinationMongodbUpdateAuthorizationTypeLoginPasswordAuthorization = "login/password"
)

func (e DestinationMongodbUpdateAuthorizationTypeLoginPasswordAuthorization) ToPointer() *DestinationMongodbUpdateAuthorizationTypeLoginPasswordAuthorization {
	return &e
}

func (e *DestinationMongodbUpdateAuthorizationTypeLoginPasswordAuthorization) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "login/password":
		*e = DestinationMongodbUpdateAuthorizationTypeLoginPasswordAuthorization(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateAuthorizationTypeLoginPasswordAuthorization: %v", v)
	}
}

// DestinationMongodbUpdateAuthorizationTypeLoginPassword - Login/Password.
type DestinationMongodbUpdateAuthorizationTypeLoginPassword struct {
	Authorization DestinationMongodbUpdateAuthorizationTypeLoginPasswordAuthorization `json:"authorization"`
	// Password associated with the username.
	Password string `json:"password"`
	// Username to use to access the database.
	Username string `json:"username"`
}

type DestinationMongodbUpdateAuthorizationTypeNoneAuthorization string

const (
	DestinationMongodbUpdateAuthorizationTypeNoneAuthorizationNone DestinationMongodbUpdateAuthorizationTypeNoneAuthorization = "none"
)

func (e DestinationMongodbUpdateAuthorizationTypeNoneAuthorization) ToPointer() *DestinationMongodbUpdateAuthorizationTypeNoneAuthorization {
	return &e
}

func (e *DestinationMongodbUpdateAuthorizationTypeNoneAuthorization) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		*e = DestinationMongodbUpdateAuthorizationTypeNoneAuthorization(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateAuthorizationTypeNoneAuthorization: %v", v)
	}
}

// DestinationMongodbUpdateAuthorizationTypeNone - None.
type DestinationMongodbUpdateAuthorizationTypeNone struct {
	Authorization DestinationMongodbUpdateAuthorizationTypeNoneAuthorization `json:"authorization"`
}

type DestinationMongodbUpdateAuthorizationTypeType string

const (
	DestinationMongodbUpdateAuthorizationTypeTypeDestinationMongodbUpdateAuthorizationTypeNone          DestinationMongodbUpdateAuthorizationTypeType = "destination-mongodb-update_Authorization type_None"
	DestinationMongodbUpdateAuthorizationTypeTypeDestinationMongodbUpdateAuthorizationTypeLoginPassword DestinationMongodbUpdateAuthorizationTypeType = "destination-mongodb-update_Authorization type_Login/Password"
)

type DestinationMongodbUpdateAuthorizationType struct {
	DestinationMongodbUpdateAuthorizationTypeNone          *DestinationMongodbUpdateAuthorizationTypeNone
	DestinationMongodbUpdateAuthorizationTypeLoginPassword *DestinationMongodbUpdateAuthorizationTypeLoginPassword

	Type DestinationMongodbUpdateAuthorizationTypeType
}

func CreateDestinationMongodbUpdateAuthorizationTypeDestinationMongodbUpdateAuthorizationTypeNone(destinationMongodbUpdateAuthorizationTypeNone DestinationMongodbUpdateAuthorizationTypeNone) DestinationMongodbUpdateAuthorizationType {
	typ := DestinationMongodbUpdateAuthorizationTypeTypeDestinationMongodbUpdateAuthorizationTypeNone

	return DestinationMongodbUpdateAuthorizationType{
		DestinationMongodbUpdateAuthorizationTypeNone: &destinationMongodbUpdateAuthorizationTypeNone,
		Type: typ,
	}
}

func CreateDestinationMongodbUpdateAuthorizationTypeDestinationMongodbUpdateAuthorizationTypeLoginPassword(destinationMongodbUpdateAuthorizationTypeLoginPassword DestinationMongodbUpdateAuthorizationTypeLoginPassword) DestinationMongodbUpdateAuthorizationType {
	typ := DestinationMongodbUpdateAuthorizationTypeTypeDestinationMongodbUpdateAuthorizationTypeLoginPassword

	return DestinationMongodbUpdateAuthorizationType{
		DestinationMongodbUpdateAuthorizationTypeLoginPassword: &destinationMongodbUpdateAuthorizationTypeLoginPassword,
		Type: typ,
	}
}

func (u *DestinationMongodbUpdateAuthorizationType) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMongodbUpdateAuthorizationTypeNone := new(DestinationMongodbUpdateAuthorizationTypeNone)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbUpdateAuthorizationTypeNone); err == nil {
		u.DestinationMongodbUpdateAuthorizationTypeNone = destinationMongodbUpdateAuthorizationTypeNone
		u.Type = DestinationMongodbUpdateAuthorizationTypeTypeDestinationMongodbUpdateAuthorizationTypeNone
		return nil
	}

	destinationMongodbUpdateAuthorizationTypeLoginPassword := new(DestinationMongodbUpdateAuthorizationTypeLoginPassword)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbUpdateAuthorizationTypeLoginPassword); err == nil {
		u.DestinationMongodbUpdateAuthorizationTypeLoginPassword = destinationMongodbUpdateAuthorizationTypeLoginPassword
		u.Type = DestinationMongodbUpdateAuthorizationTypeTypeDestinationMongodbUpdateAuthorizationTypeLoginPassword
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMongodbUpdateAuthorizationType) MarshalJSON() ([]byte, error) {
	if u.DestinationMongodbUpdateAuthorizationTypeNone != nil {
		return json.Marshal(u.DestinationMongodbUpdateAuthorizationTypeNone)
	}

	if u.DestinationMongodbUpdateAuthorizationTypeLoginPassword != nil {
		return json.Marshal(u.DestinationMongodbUpdateAuthorizationTypeLoginPassword)
	}

	return nil, nil
}

type DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstance string

const (
	DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstanceAtlas DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstance = "atlas"
)

func (e DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstance) ToPointer() *DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstance {
	return &e
}

func (e *DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "atlas":
		*e = DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstance: %v", v)
	}
}

// DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas struct {
	// URL of a cluster to connect to.
	ClusterURL string                                                          `json:"cluster_url"`
	Instance   DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstance `json:"instance"`
}

type DestinationMongodbUpdateMongoDbInstanceTypeReplicaSetInstance string

const (
	DestinationMongodbUpdateMongoDbInstanceTypeReplicaSetInstanceReplica DestinationMongodbUpdateMongoDbInstanceTypeReplicaSetInstance = "replica"
)

func (e DestinationMongodbUpdateMongoDbInstanceTypeReplicaSetInstance) ToPointer() *DestinationMongodbUpdateMongoDbInstanceTypeReplicaSetInstance {
	return &e
}

func (e *DestinationMongodbUpdateMongoDbInstanceTypeReplicaSetInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "replica":
		*e = DestinationMongodbUpdateMongoDbInstanceTypeReplicaSetInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateMongoDbInstanceTypeReplicaSetInstance: %v", v)
	}
}

// DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet struct {
	Instance DestinationMongodbUpdateMongoDbInstanceTypeReplicaSetInstance `json:"instance"`
	// A replica set name.
	ReplicaSet *string `json:"replica_set,omitempty"`
	// The members of a replica set. Please specify `host`:`port` of each member seperated by comma.
	ServerAddresses string `json:"server_addresses"`
}

type DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstance string

const (
	DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceStandalone DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstance = "standalone"
)

func (e DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstance) ToPointer() *DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstance {
	return &e
}

func (e *DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standalone":
		*e = DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstance: %v", v)
	}
}

// DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance struct {
	// The Host of a Mongo database to be replicated.
	Host     string                                                                       `json:"host"`
	Instance DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstance `json:"instance"`
	// The Port of a Mongo database to be replicated.
	Port int64 `json:"port"`
}

type DestinationMongodbUpdateMongoDbInstanceTypeType string

const (
	DestinationMongodbUpdateMongoDbInstanceTypeTypeDestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance DestinationMongodbUpdateMongoDbInstanceTypeType = "destination-mongodb-update_MongoDb Instance Type_Standalone MongoDb Instance"
	DestinationMongodbUpdateMongoDbInstanceTypeTypeDestinationMongodbUpdateMongoDbInstanceTypeReplicaSet                DestinationMongodbUpdateMongoDbInstanceTypeType = "destination-mongodb-update_MongoDb Instance Type_Replica Set"
	DestinationMongodbUpdateMongoDbInstanceTypeTypeDestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas              DestinationMongodbUpdateMongoDbInstanceTypeType = "destination-mongodb-update_MongoDb Instance Type_MongoDB Atlas"
)

type DestinationMongodbUpdateMongoDbInstanceType struct {
	DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance *DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance
	DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet                *DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet
	DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas              *DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas

	Type DestinationMongodbUpdateMongoDbInstanceTypeType
}

func CreateDestinationMongodbUpdateMongoDbInstanceTypeDestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance(destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance) DestinationMongodbUpdateMongoDbInstanceType {
	typ := DestinationMongodbUpdateMongoDbInstanceTypeTypeDestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance

	return DestinationMongodbUpdateMongoDbInstanceType{
		DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance: &destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance,
		Type: typ,
	}
}

func CreateDestinationMongodbUpdateMongoDbInstanceTypeDestinationMongodbUpdateMongoDbInstanceTypeReplicaSet(destinationMongodbUpdateMongoDbInstanceTypeReplicaSet DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet) DestinationMongodbUpdateMongoDbInstanceType {
	typ := DestinationMongodbUpdateMongoDbInstanceTypeTypeDestinationMongodbUpdateMongoDbInstanceTypeReplicaSet

	return DestinationMongodbUpdateMongoDbInstanceType{
		DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet: &destinationMongodbUpdateMongoDbInstanceTypeReplicaSet,
		Type: typ,
	}
}

func CreateDestinationMongodbUpdateMongoDbInstanceTypeDestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas(destinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas) DestinationMongodbUpdateMongoDbInstanceType {
	typ := DestinationMongodbUpdateMongoDbInstanceTypeTypeDestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas

	return DestinationMongodbUpdateMongoDbInstanceType{
		DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas: &destinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas,
		Type: typ,
	}
}

func (u *DestinationMongodbUpdateMongoDbInstanceType) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas := new(DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas); err == nil {
		u.DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas = destinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas
		u.Type = DestinationMongodbUpdateMongoDbInstanceTypeTypeDestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas
		return nil
	}

	destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance := new(DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance); err == nil {
		u.DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance = destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance
		u.Type = DestinationMongodbUpdateMongoDbInstanceTypeTypeDestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance
		return nil
	}

	destinationMongodbUpdateMongoDbInstanceTypeReplicaSet := new(DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbUpdateMongoDbInstanceTypeReplicaSet); err == nil {
		u.DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet = destinationMongodbUpdateMongoDbInstanceTypeReplicaSet
		u.Type = DestinationMongodbUpdateMongoDbInstanceTypeTypeDestinationMongodbUpdateMongoDbInstanceTypeReplicaSet
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMongodbUpdateMongoDbInstanceType) MarshalJSON() ([]byte, error) {
	if u.DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas != nil {
		return json.Marshal(u.DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas)
	}

	if u.DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
		return json.Marshal(u.DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance)
	}

	if u.DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet != nil {
		return json.Marshal(u.DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet)
	}

	return nil, nil
}

// DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationMongodbUpdateSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbUpdateSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethod `json:"tunnel_method"`
}

type DestinationMongodbUpdateSSHTunnelMethodType string

const (
	DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHTunnelMethodNoTunnel               DestinationMongodbUpdateSSHTunnelMethodType = "destination-mongodb-update_SSH Tunnel Method_No Tunnel"
	DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication   DestinationMongodbUpdateSSHTunnelMethodType = "destination-mongodb-update_SSH Tunnel Method_SSH Key Authentication"
	DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication DestinationMongodbUpdateSSHTunnelMethodType = "destination-mongodb-update_SSH Tunnel Method_Password Authentication"
)

type DestinationMongodbUpdateSSHTunnelMethod struct {
	DestinationMongodbUpdateSSHTunnelMethodNoTunnel               *DestinationMongodbUpdateSSHTunnelMethodNoTunnel
	DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication   *DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication
	DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication *DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication

	Type DestinationMongodbUpdateSSHTunnelMethodType
}

func CreateDestinationMongodbUpdateSSHTunnelMethodDestinationMongodbUpdateSSHTunnelMethodNoTunnel(destinationMongodbUpdateSSHTunnelMethodNoTunnel DestinationMongodbUpdateSSHTunnelMethodNoTunnel) DestinationMongodbUpdateSSHTunnelMethod {
	typ := DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHTunnelMethodNoTunnel

	return DestinationMongodbUpdateSSHTunnelMethod{
		DestinationMongodbUpdateSSHTunnelMethodNoTunnel: &destinationMongodbUpdateSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationMongodbUpdateSSHTunnelMethodDestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication(destinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication) DestinationMongodbUpdateSSHTunnelMethod {
	typ := DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication

	return DestinationMongodbUpdateSSHTunnelMethod{
		DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication: &destinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationMongodbUpdateSSHTunnelMethodDestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication(destinationMongodbUpdateSSHTunnelMethodPasswordAuthentication DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication) DestinationMongodbUpdateSSHTunnelMethod {
	typ := DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication

	return DestinationMongodbUpdateSSHTunnelMethod{
		DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication: &destinationMongodbUpdateSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationMongodbUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMongodbUpdateSSHTunnelMethodNoTunnel := new(DestinationMongodbUpdateSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbUpdateSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationMongodbUpdateSSHTunnelMethodNoTunnel = destinationMongodbUpdateSSHTunnelMethodNoTunnel
		u.Type = DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHTunnelMethodNoTunnel
		return nil
	}

	destinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication := new(DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication = destinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationMongodbUpdateSSHTunnelMethodPasswordAuthentication := new(DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbUpdateSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication = destinationMongodbUpdateSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMongodbUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMongodbUpdateSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationMongodbUpdateSSHTunnelMethodNoTunnel)
	}

	if u.DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type DestinationMongodbUpdate struct {
	// Authorization type.
	AuthType DestinationMongodbUpdateAuthorizationType `json:"auth_type"`
	// Name of the database.
	Database string `json:"database"`
	// MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
	InstanceType *DestinationMongodbUpdateMongoDbInstanceType `json:"instance_type,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMongodbUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
}
