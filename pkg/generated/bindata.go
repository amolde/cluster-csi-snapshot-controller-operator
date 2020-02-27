// Code generated by go-bindata.
// sources:
// assets/csi_controller_deployment.yaml
// assets/volumesnapshotclasses.yaml
// assets/volumesnapshotcontents.yaml
// assets/volumesnapshots.yaml
// DO NOT EDIT!

package generated

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _csi_controller_deploymentYaml = []byte(`kind: Deployment
apiVersion: apps/v1
metadata:
  name: csi-snapshot-controller
  namespace: openshift-csi-snapshot-controller
spec:
  serviceName: "csi-snapshot-controller"
  replicas: 1
  selector:
    matchLabels:
      app: csi-snapshot-controller
  template:
    metadata:
      labels:
        app: csi-snapshot-controller
    spec:
      serviceAccount: csi-snapshot-controller
      containers:
        - name: snapshot-controller
          image: quay.io/k8scsi/snapshot-controller:v2.0.0
          args:
            - "--v=5"
            - "--leader-election=true"
          imagePullPolicy: IfNotPresent
          resources:
            requests:
              # TODO: measure on a real cluster
              cpu: 10m
              memory: 50Mi
      priorityClassName: "system-cluster-critical"
      tolerations:
      - key: "node.kubernetes.io/unreachable"
        operator: "Exists"
        effect: "NoExecute"
        tolerationSeconds: 120
      - key: "node.kubernetes.io/not-ready"
        operator: "Exists"
        effect: "NoExecute"
        tolerationSeconds: 120
`)

func csi_controller_deploymentYamlBytes() ([]byte, error) {
	return _csi_controller_deploymentYaml, nil
}

func csi_controller_deploymentYaml() (*asset, error) {
	bytes, err := csi_controller_deploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "csi_controller_deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _volumesnapshotclassesYaml = []byte(`
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.5
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/260"
  creationTimestamp: null
  name: volumesnapshotclasses.snapshot.storage.k8s.io
spec:
  additionalPrinterColumns:
  - JSONPath: .driver
    name: Driver
    type: string
  - JSONPath: .deletionPolicy
    description: Determines whether a VolumeSnapshotContent created through the VolumeSnapshotClass
      should be deleted when its bound VolumeSnapshot is deleted.
    name: DeletionPolicy
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: snapshot.storage.k8s.io
  names:
    kind: VolumeSnapshotClass
    listKind: VolumeSnapshotClassList
    plural: volumesnapshotclasses
    singular: volumesnapshotclass
  preserveUnknownFields: false
  scope: Cluster
  subresources: {}
  validation:
    openAPIV3Schema:
      description: VolumeSnapshotClass specifies parameters that a underlying storage
        system uses when creating a volume snapshot. A specific VolumeSnapshotClass
        is used by specifying its name in a VolumeSnapshot object. VolumeSnapshotClasses
        are non-namespaced
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        deletionPolicy:
          description: deletionPolicy determines whether a VolumeSnapshotContent created
            through the VolumeSnapshotClass should be deleted when its bound VolumeSnapshot
            is deleted. Supported values are "Retain" and "Delete". "Retain" means
            that the VolumeSnapshotContent and its physical snapshot on underlying
            storage system are kept. "Delete" means that the VolumeSnapshotContent
            and its physical snapshot on underlying storage system are deleted. Required.
          enum:
          - Delete
          - Retain
          type: string
        driver:
          description: driver is the name of the storage driver that handles this
            VolumeSnapshotClass. Required.
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        parameters:
          additionalProperties:
            type: string
          description: parameters is a key-value map with storage driver specific
            parameters for creating snapshots. These values are opaque to Kubernetes.
          type: object
      required:
      - deletionPolicy
      - driver
      type: object
  version: v1beta1
  versions:
  - name: v1beta1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func volumesnapshotclassesYamlBytes() ([]byte, error) {
	return _volumesnapshotclassesYaml, nil
}

func volumesnapshotclassesYaml() (*asset, error) {
	bytes, err := volumesnapshotclassesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "volumesnapshotclasses.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _volumesnapshotcontentsYaml = []byte(`
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.5
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/260"
  creationTimestamp: null
  name: volumesnapshotcontents.snapshot.storage.k8s.io
spec:
  additionalPrinterColumns:
  - JSONPath: .status.readyToUse
    description: Indicates if a snapshot is ready to be used to restore a volume.
    name: ReadyToUse
    type: boolean
  - JSONPath: .status.restoreSize
    description: Represents the complete size of the snapshot in bytes
    name: RestoreSize
    type: integer
  - JSONPath: .spec.deletionPolicy
    description: Determines whether this VolumeSnapshotContent and its physical snapshot
      on the underlying storage system should be deleted when its bound VolumeSnapshot
      is deleted.
    name: DeletionPolicy
    type: string
  - JSONPath: .spec.driver
    description: Name of the CSI driver used to create the physical snapshot on the
      underlying storage system.
    name: Driver
    type: string
  - JSONPath: .spec.volumeSnapshotClassName
    description: Name of the VolumeSnapshotClass to which this snapshot belongs.
    name: VolumeSnapshotClass
    type: string
  - JSONPath: .spec.volumeSnapshotRef.name
    description: Name of the VolumeSnapshot object to which this VolumeSnapshotContent
      object is bound.
    name: VolumeSnapshot
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: snapshot.storage.k8s.io
  names:
    kind: VolumeSnapshotContent
    listKind: VolumeSnapshotContentList
    plural: volumesnapshotcontents
    singular: volumesnapshotcontent
  preserveUnknownFields: false
  scope: Cluster
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: VolumeSnapshotContent represents the actual "on-disk" snapshot
        object in the underlying storage system
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        spec:
          description: spec defines properties of a VolumeSnapshotContent created
            by the underlying storage system. Required.
          properties:
            deletionPolicy:
              description: deletionPolicy determines whether this VolumeSnapshotContent
                and its physical snapshot on the underlying storage system should
                be deleted when its bound VolumeSnapshot is deleted. Supported values
                are "Retain" and "Delete". "Retain" means that the VolumeSnapshotContent
                and its physical snapshot on underlying storage system are kept. "Delete"
                means that the VolumeSnapshotContent and its physical snapshot on
                underlying storage system are deleted. In dynamic snapshot creation
                case, this field will be filled in with the "DeletionPolicy" field
                defined in the VolumeSnapshotClass the VolumeSnapshot refers to. For
                pre-existing snapshots, users MUST specify this field when creating
                the VolumeSnapshotContent object. Required.
              enum:
              - Delete
              - Retain
              type: string
            driver:
              description: driver is the name of the CSI driver used to create the
                physical snapshot on the underlying storage system. This MUST be the
                same as the name returned by the CSI GetPluginName() call for that
                driver. Required.
              type: string
            source:
              description: source specifies from where a snapshot will be created.
                This field is immutable after creation. Required.
              properties:
                snapshotHandle:
                  description: snapshotHandle specifies the CSI "snapshot_id" of a
                    pre-existing snapshot on the underlying storage system. This field
                    is immutable.
                  type: string
                volumeHandle:
                  description: volumeHandle specifies the CSI "volume_id" of the volume
                    from which a snapshot should be dynamically taken from. This field
                    is immutable.
                  type: string
              type: object
            volumeSnapshotClassName:
              description: name of the VolumeSnapshotClass to which this snapshot
                belongs.
              type: string
            volumeSnapshotRef:
              description: volumeSnapshotRef specifies the VolumeSnapshot object to
                which this VolumeSnapshotContent object is bound. VolumeSnapshot.Spec.VolumeSnapshotContentName
                field must reference to this VolumeSnapshotContent's name for the
                bidirectional binding to be valid. For a pre-existing VolumeSnapshotContent
                object, name and namespace of the VolumeSnapshot object MUST be provided
                for binding to happen. This field is immutable after creation. Required.
              properties:
                apiVersion:
                  description: API version of the referent.
                  type: string
                fieldPath:
                  description: 'If referring to a piece of an object instead of an
                    entire object, this string should contain a valid JSON/Go field
                    access statement, such as desiredState.manifest.containers[2].
                    For example, if the object reference is to a container within
                    a pod, this would take on a value like: "spec.containers{name}"
                    (where "name" refers to the name of the container that triggered
                    the event) or if no container name is specified "spec.containers[2]"
                    (container with index 2 in this pod). This syntax is chosen only
                    to have some well-defined way of referencing a part of an object.
                    TODO: this design is not final and this field is subject to change
                    in the future.'
                  type: string
                kind:
                  description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
                resourceVersion:
                  description: 'Specific resourceVersion to which this reference is
                    made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                  type: string
                uid:
                  description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                  type: string
              type: object
          required:
          - deletionPolicy
          - driver
          - source
          - volumeSnapshotRef
          type: object
        status:
          description: status represents the current information of a snapshot.
          properties:
            creationTime:
              description: creationTime is the timestamp when the point-in-time snapshot
                is taken by the underlying storage system. In dynamic snapshot creation
                case, this field will be filled in with the "creation_time" value
                returned from CSI "CreateSnapshotRequest" gRPC call. For a pre-existing
                snapshot, this field will be filled with the "creation_time" value
                returned from the CSI "ListSnapshots" gRPC call if the driver supports
                it. If not specified, it indicates the creation time is unknown. The
                format of this field is a Unix nanoseconds time encoded as an int64.
                On Unix, the command ` + "`" + `date +%s%N` + "`" + ` returns the current time in nanoseconds
                since 1970-01-01 00:00:00 UTC.
              format: int64
              type: integer
            error:
              description: error is the latest observed error during snapshot creation,
                if any.
              properties:
                message:
                  description: 'message is a string detailing the encountered error
                    during snapshot creation if specified. NOTE: message may be logged,
                    and it should not contain sensitive information.'
                  type: string
                time:
                  description: time is the timestamp when the error was encountered.
                  format: date-time
                  type: string
              type: object
            readyToUse:
              description: readyToUse indicates if a snapshot is ready to be used
                to restore a volume. In dynamic snapshot creation case, this field
                will be filled in with the "ready_to_use" value returned from CSI
                "CreateSnapshotRequest" gRPC call. For a pre-existing snapshot, this
                field will be filled with the "ready_to_use" value returned from the
                CSI "ListSnapshots" gRPC call if the driver supports it, otherwise,
                this field will be set to "True". If not specified, it means the readiness
                of a snapshot is unknown.
              type: boolean
            restoreSize:
              description: restoreSize represents the complete size of the snapshot
                in bytes. In dynamic snapshot creation case, this field will be filled
                in with the "size_bytes" value returned from CSI "CreateSnapshotRequest"
                gRPC call. For a pre-existing snapshot, this field will be filled
                with the "size_bytes" value returned from the CSI "ListSnapshots"
                gRPC call if the driver supports it. When restoring a volume from
                this snapshot, the size of the volume MUST NOT be smaller than the
                restoreSize if it is specified, otherwise the restoration will fail.
                If not specified, it indicates that the size is unknown.
              format: int64
              minimum: 0
              type: integer
            snapshotHandle:
              description: snapshotHandle is the CSI "snapshot_id" of a snapshot on
                the underlying storage system. If not specified, it indicates that
                dynamic snapshot creation has either failed or it is still in progress.
              type: string
          type: object
      required:
      - spec
      type: object
  version: v1beta1
  versions:
  - name: v1beta1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func volumesnapshotcontentsYamlBytes() ([]byte, error) {
	return _volumesnapshotcontentsYaml, nil
}

func volumesnapshotcontentsYaml() (*asset, error) {
	bytes, err := volumesnapshotcontentsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "volumesnapshotcontents.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _volumesnapshotsYaml = []byte(`
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.5
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/260"
  creationTimestamp: null
  name: volumesnapshots.snapshot.storage.k8s.io
spec:
  additionalPrinterColumns:
  - JSONPath: .status.readyToUse
    description: Indicates if a snapshot is ready to be used to restore a volume.
    name: ReadyToUse
    type: boolean
  - JSONPath: .spec.source.persistentVolumeClaimName
    description: Name of the source PVC from where a dynamically taken snapshot will
      be created.
    name: SourcePVC
    type: string
  - JSONPath: .spec.source.volumeSnapshotContentName
    description: Name of the VolumeSnapshotContent which represents a pre-provisioned
      snapshot.
    name: SourceSnapshotContent
    type: string
  - JSONPath: .status.restoreSize
    description: Represents the complete size of the snapshot.
    name: RestoreSize
    type: string
  - JSONPath: .spec.volumeSnapshotClassName
    description: The name of the VolumeSnapshotClass requested by the VolumeSnapshot.
    name: SnapshotClass
    type: string
  - JSONPath: .status.boundVolumeSnapshotContentName
    description: The name of the VolumeSnapshotContent to which this VolumeSnapshot
      is bound.
    name: SnapshotContent
    type: string
  - JSONPath: .status.creationTime
    description: Timestamp when the point-in-time snapshot is taken by the underlying
      storage system.
    name: CreationTime
    type: date
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: snapshot.storage.k8s.io
  names:
    kind: VolumeSnapshot
    listKind: VolumeSnapshotList
    plural: volumesnapshots
    singular: volumesnapshot
  preserveUnknownFields: false
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: VolumeSnapshot is a user's request for either creating a point-in-time
        snapshot of a persistent volume, or binding to a pre-existing snapshot.
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        spec:
          description: 'spec defines the desired characteristics of a snapshot requested
            by a user. More info: https://kubernetes.io/docs/concepts/storage/volume-snapshots#volumesnapshots
            Required.'
          properties:
            source:
              description: source specifies where a snapshot will be created from.
                This field is immutable after creation. Required.
              properties:
                persistentVolumeClaimName:
                  description: persistentVolumeClaimName specifies the name of the
                    PersistentVolumeClaim object in the same namespace as the VolumeSnapshot
                    object where the snapshot should be dynamically taken from. This
                    field is immutable.
                  type: string
                volumeSnapshotContentName:
                  description: volumeSnapshotContentName specifies the name of a pre-existing
                    VolumeSnapshotContent object. This field is immutable.
                  type: string
              type: object
            volumeSnapshotClassName:
              description: 'volumeSnapshotClassName is the name of the VolumeSnapshotClass
                requested by the VolumeSnapshot. If not specified, the default snapshot
                class will be used if one exists. If not specified, and there is no
                default snapshot class, dynamic snapshot creation will fail. Empty
                string is not allowed for this field. TODO(xiangqian): a webhook validation
                on empty string. More info: https://kubernetes.io/docs/concepts/storage/volume-snapshot-classes'
              type: string
          required:
          - source
          type: object
        status:
          description: 'status represents the current information of a snapshot. NOTE:
            status can be modified by sources other than system controllers, and must
            not be depended upon for accuracy. Controllers should only use information
            from the VolumeSnapshotContent object after verifying that the binding
            is accurate and complete.'
          properties:
            boundVolumeSnapshotContentName:
              description: 'boundVolumeSnapshotContentName represents the name of
                the VolumeSnapshotContent object to which the VolumeSnapshot object
                is bound. If not specified, it indicates that the VolumeSnapshot object
                has not been successfully bound to a VolumeSnapshotContent object
                yet. NOTE: Specified boundVolumeSnapshotContentName alone does not
                mean binding       is valid. Controllers MUST always verify bidirectional
                binding between       VolumeSnapshot and VolumeSnapshotContent to
                avoid possible security issues.'
              type: string
            creationTime:
              description: creationTime is the timestamp when the point-in-time snapshot
                is taken by the underlying storage system. In dynamic snapshot creation
                case, this field will be filled in with the "creation_time" value
                returned from CSI "CreateSnapshotRequest" gRPC call. For a pre-existing
                snapshot, this field will be filled with the "creation_time" value
                returned from the CSI "ListSnapshots" gRPC call if the driver supports
                it. If not specified, it indicates that the creation time of the snapshot
                is unknown.
              format: date-time
              type: string
            error:
              description: error is the last observed error during snapshot creation,
                if any. This field could be helpful to upper level controllers(i.e.,
                application controller) to decide whether they should continue on
                waiting for the snapshot to be created based on the type of error
                reported.
              properties:
                message:
                  description: 'message is a string detailing the encountered error
                    during snapshot creation if specified. NOTE: message may be logged,
                    and it should not contain sensitive information.'
                  type: string
                time:
                  description: time is the timestamp when the error was encountered.
                  format: date-time
                  type: string
              type: object
            readyToUse:
              description: readyToUse indicates if a snapshot is ready to be used
                to restore a volume. In dynamic snapshot creation case, this field
                will be filled in with the "ready_to_use" value returned from CSI
                "CreateSnapshotRequest" gRPC call. For a pre-existing snapshot, this
                field will be filled with the "ready_to_use" value returned from the
                CSI "ListSnapshots" gRPC call if the driver supports it, otherwise,
                this field will be set to "True". If not specified, it means the readiness
                of a snapshot is unknown.
              type: boolean
            restoreSize:
              anyOf:
              - type: integer
              - type: string
              description: restoreSize represents the complete size of the snapshot
                in bytes. In dynamic snapshot creation case, this field will be filled
                in with the "size_bytes" value returned from CSI "CreateSnapshotRequest"
                gRPC call. For a pre-existing snapshot, this field will be filled
                with the "size_bytes" value returned from the CSI "ListSnapshots"
                gRPC call if the driver supports it. When restoring a volume from
                this snapshot, the size of the volume MUST NOT be smaller than the
                restoreSize if it is specified, otherwise the restoration will fail.
                If not specified, it indicates that the size is unknown.
              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
              x-kubernetes-int-or-string: true
          type: object
      required:
      - spec
      type: object
  version: v1beta1
  versions:
  - name: v1beta1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func volumesnapshotsYamlBytes() ([]byte, error) {
	return _volumesnapshotsYaml, nil
}

func volumesnapshotsYaml() (*asset, error) {
	bytes, err := volumesnapshotsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "volumesnapshots.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"csi_controller_deployment.yaml": csi_controller_deploymentYaml,
	"volumesnapshotclasses.yaml":     volumesnapshotclassesYaml,
	"volumesnapshotcontents.yaml":    volumesnapshotcontentsYaml,
	"volumesnapshots.yaml":           volumesnapshotsYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"csi_controller_deployment.yaml": {csi_controller_deploymentYaml, map[string]*bintree{}},
	"volumesnapshotclasses.yaml":     {volumesnapshotclassesYaml, map[string]*bintree{}},
	"volumesnapshotcontents.yaml":    {volumesnapshotcontentsYaml, map[string]*bintree{}},
	"volumesnapshots.yaml":           {volumesnapshotsYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
