// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/containeranalysis/v1alpha1/image_basis.proto

package containeranalysis

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Instructions from dockerfile
type DockerImage_Layer_Directive int32

const (
	// Default value for unsupported/missing directive
	DockerImage_Layer_DIRECTIVE_UNSPECIFIED DockerImage_Layer_Directive = 0
	// https://docs.docker.com/reference/builder/#maintainer
	DockerImage_Layer_MAINTAINER DockerImage_Layer_Directive = 1
	// https://docs.docker.com/reference/builder/#run
	DockerImage_Layer_RUN DockerImage_Layer_Directive = 2
	// https://docs.docker.com/reference/builder/#cmd
	DockerImage_Layer_CMD DockerImage_Layer_Directive = 3
	// https://docs.docker.com/reference/builder/#label
	DockerImage_Layer_LABEL DockerImage_Layer_Directive = 4
	// https://docs.docker.com/reference/builder/#expose
	DockerImage_Layer_EXPOSE DockerImage_Layer_Directive = 5
	// https://docs.docker.com/reference/builder/#env
	DockerImage_Layer_ENV DockerImage_Layer_Directive = 6
	// https://docs.docker.com/reference/builder/#add
	DockerImage_Layer_ADD DockerImage_Layer_Directive = 7
	// https://docs.docker.com/reference/builder/#copy
	DockerImage_Layer_COPY DockerImage_Layer_Directive = 8
	// https://docs.docker.com/reference/builder/#entrypoint
	DockerImage_Layer_ENTRYPOINT DockerImage_Layer_Directive = 9
	// https://docs.docker.com/reference/builder/#volume
	DockerImage_Layer_VOLUME DockerImage_Layer_Directive = 10
	// https://docs.docker.com/reference/builder/#user
	DockerImage_Layer_USER DockerImage_Layer_Directive = 11
	// https://docs.docker.com/reference/builder/#workdir
	DockerImage_Layer_WORKDIR DockerImage_Layer_Directive = 12
	// https://docs.docker.com/reference/builder/#arg
	DockerImage_Layer_ARG DockerImage_Layer_Directive = 13
	// https://docs.docker.com/reference/builder/#onbuild
	DockerImage_Layer_ONBUILD DockerImage_Layer_Directive = 14
	// https://docs.docker.com/reference/builder/#stopsignal
	DockerImage_Layer_STOPSIGNAL DockerImage_Layer_Directive = 15
	// https://docs.docker.com/reference/builder/#healthcheck
	DockerImage_Layer_HEALTHCHECK DockerImage_Layer_Directive = 16
	// https://docs.docker.com/reference/builder/#shell
	DockerImage_Layer_SHELL DockerImage_Layer_Directive = 17
)

var DockerImage_Layer_Directive_name = map[int32]string{
	0:  "DIRECTIVE_UNSPECIFIED",
	1:  "MAINTAINER",
	2:  "RUN",
	3:  "CMD",
	4:  "LABEL",
	5:  "EXPOSE",
	6:  "ENV",
	7:  "ADD",
	8:  "COPY",
	9:  "ENTRYPOINT",
	10: "VOLUME",
	11: "USER",
	12: "WORKDIR",
	13: "ARG",
	14: "ONBUILD",
	15: "STOPSIGNAL",
	16: "HEALTHCHECK",
	17: "SHELL",
}
var DockerImage_Layer_Directive_value = map[string]int32{
	"DIRECTIVE_UNSPECIFIED": 0,
	"MAINTAINER":            1,
	"RUN":                   2,
	"CMD":                   3,
	"LABEL":                 4,
	"EXPOSE":                5,
	"ENV":                   6,
	"ADD":                   7,
	"COPY":                  8,
	"ENTRYPOINT":            9,
	"VOLUME":                10,
	"USER":                  11,
	"WORKDIR":               12,
	"ARG":                   13,
	"ONBUILD":               14,
	"STOPSIGNAL":            15,
	"HEALTHCHECK":           16,
	"SHELL":                 17,
}

func (x DockerImage_Layer_Directive) String() string {
	return proto.EnumName(DockerImage_Layer_Directive_name, int32(x))
}
func (DockerImage_Layer_Directive) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 0, 0}
}

// DockerImage holds types defining base image notes
// and derived image occurrences.
type DockerImage struct {
}

func (m *DockerImage) Reset()                    { *m = DockerImage{} }
func (m *DockerImage) String() string            { return proto.CompactTextString(m) }
func (*DockerImage) ProtoMessage()               {}
func (*DockerImage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// Layer holds metadata specific to a layer of a Docker image.
type DockerImage_Layer struct {
	// The recovered Dockerfile directive used to construct this layer.
	Directive DockerImage_Layer_Directive `protobuf:"varint,1,opt,name=directive,enum=google.devtools.containeranalysis.v1alpha1.DockerImage_Layer_Directive" json:"directive,omitempty"`
	// The recovered arguments to the Dockerfile directive.
	Arguments string `protobuf:"bytes,2,opt,name=arguments" json:"arguments,omitempty"`
}

func (m *DockerImage_Layer) Reset()                    { *m = DockerImage_Layer{} }
func (m *DockerImage_Layer) String() string            { return proto.CompactTextString(m) }
func (*DockerImage_Layer) ProtoMessage()               {}
func (*DockerImage_Layer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *DockerImage_Layer) GetDirective() DockerImage_Layer_Directive {
	if m != nil {
		return m.Directive
	}
	return DockerImage_Layer_DIRECTIVE_UNSPECIFIED
}

func (m *DockerImage_Layer) GetArguments() string {
	if m != nil {
		return m.Arguments
	}
	return ""
}

// A set of properties that uniquely identify a given Docker image.
type DockerImage_Fingerprint struct {
	// The layer-id of the final layer in the Docker image's v1
	// representation.
	// This field can be used as a filter in list requests.
	V1Name string `protobuf:"bytes,1,opt,name=v1_name,json=v1Name" json:"v1_name,omitempty"`
	// The ordered list of v2 blobs that represent a given image.
	V2Blob []string `protobuf:"bytes,2,rep,name=v2_blob,json=v2Blob" json:"v2_blob,omitempty"`
	// Output only. The name of the image's v2 blobs computed via:
	//   [bottom] := v2_blob[bottom]
	//   [N] := sha256(v2_blob[N] + " " + v2_name[N+1])
	// Only the name of the final blob is kept.
	// This field can be used as a filter in list requests.
	V2Name string `protobuf:"bytes,3,opt,name=v2_name,json=v2Name" json:"v2_name,omitempty"`
}

func (m *DockerImage_Fingerprint) Reset()                    { *m = DockerImage_Fingerprint{} }
func (m *DockerImage_Fingerprint) String() string            { return proto.CompactTextString(m) }
func (*DockerImage_Fingerprint) ProtoMessage()               {}
func (*DockerImage_Fingerprint) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 1} }

func (m *DockerImage_Fingerprint) GetV1Name() string {
	if m != nil {
		return m.V1Name
	}
	return ""
}

func (m *DockerImage_Fingerprint) GetV2Blob() []string {
	if m != nil {
		return m.V2Blob
	}
	return nil
}

func (m *DockerImage_Fingerprint) GetV2Name() string {
	if m != nil {
		return m.V2Name
	}
	return ""
}

// Basis describes the base image portion (Note) of the DockerImage
// relationship.  Linked occurrences are derived from this or an
// equivalent image via:
//   FROM <Basis.resource_url>
// Or an equivalent reference, e.g. a tag of the resource_url.
type DockerImage_Basis struct {
	// The resource_url for the resource representing the basis of
	// associated occurrence images.
	ResourceUrl string `protobuf:"bytes,1,opt,name=resource_url,json=resourceUrl" json:"resource_url,omitempty"`
	// The fingerprint of the base image.
	Fingerprint *DockerImage_Fingerprint `protobuf:"bytes,2,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *DockerImage_Basis) Reset()                    { *m = DockerImage_Basis{} }
func (m *DockerImage_Basis) String() string            { return proto.CompactTextString(m) }
func (*DockerImage_Basis) ProtoMessage()               {}
func (*DockerImage_Basis) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 2} }

func (m *DockerImage_Basis) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *DockerImage_Basis) GetFingerprint() *DockerImage_Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

// Derived describes the derived image portion (Occurrence) of the
// DockerImage relationship.  This image would be produced from a Dockerfile
// with FROM <DockerImage.Basis in attached Note>.
type DockerImage_Derived struct {
	// The fingerprint of the derived image.
	Fingerprint *DockerImage_Fingerprint `protobuf:"bytes,1,opt,name=fingerprint" json:"fingerprint,omitempty"`
	// Output only. The number of layers by which this image differs from the
	// associated image basis.
	Distance uint32 `protobuf:"varint,2,opt,name=distance" json:"distance,omitempty"`
	// This contains layer-specific metadata, if populated it has length
	// "distance" and is ordered with [distance] being the layer immediately
	// following the base image and [1] being the final layer.
	LayerInfo []*DockerImage_Layer `protobuf:"bytes,3,rep,name=layer_info,json=layerInfo" json:"layer_info,omitempty"`
	// Output only. This contains the base image URL for the derived image
	// occurrence.
	BaseResourceUrl string `protobuf:"bytes,4,opt,name=base_resource_url,json=baseResourceUrl" json:"base_resource_url,omitempty"`
}

func (m *DockerImage_Derived) Reset()                    { *m = DockerImage_Derived{} }
func (m *DockerImage_Derived) String() string            { return proto.CompactTextString(m) }
func (*DockerImage_Derived) ProtoMessage()               {}
func (*DockerImage_Derived) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 3} }

func (m *DockerImage_Derived) GetFingerprint() *DockerImage_Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

func (m *DockerImage_Derived) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *DockerImage_Derived) GetLayerInfo() []*DockerImage_Layer {
	if m != nil {
		return m.LayerInfo
	}
	return nil
}

func (m *DockerImage_Derived) GetBaseResourceUrl() string {
	if m != nil {
		return m.BaseResourceUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*DockerImage)(nil), "google.devtools.containeranalysis.v1alpha1.DockerImage")
	proto.RegisterType((*DockerImage_Layer)(nil), "google.devtools.containeranalysis.v1alpha1.DockerImage.Layer")
	proto.RegisterType((*DockerImage_Fingerprint)(nil), "google.devtools.containeranalysis.v1alpha1.DockerImage.Fingerprint")
	proto.RegisterType((*DockerImage_Basis)(nil), "google.devtools.containeranalysis.v1alpha1.DockerImage.Basis")
	proto.RegisterType((*DockerImage_Derived)(nil), "google.devtools.containeranalysis.v1alpha1.DockerImage.Derived")
	proto.RegisterEnum("google.devtools.containeranalysis.v1alpha1.DockerImage_Layer_Directive", DockerImage_Layer_Directive_name, DockerImage_Layer_Directive_value)
}

func init() {
	proto.RegisterFile("google/devtools/containeranalysis/v1alpha1/image_basis.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdf, 0x6e, 0xda, 0x30,
	0x14, 0xc6, 0x17, 0x28, 0xd0, 0x9c, 0xf4, 0x8f, 0x6b, 0x69, 0x1a, 0x43, 0xbd, 0x60, 0x95, 0x26,
	0x55, 0xbd, 0x08, 0x82, 0x5d, 0x6e, 0xbb, 0x80, 0xc4, 0x85, 0xa8, 0x69, 0x40, 0x06, 0xba, 0x76,
	0x9b, 0x84, 0x0c, 0xb8, 0x59, 0xb4, 0x60, 0x23, 0x27, 0x45, 0xea, 0x3b, 0xec, 0x66, 0x37, 0x7d,
	0x80, 0x3d, 0xe1, 0xde, 0x60, 0x93, 0x53, 0x28, 0xdd, 0xaa, 0x49, 0xd5, 0xa6, 0xdd, 0x99, 0xf3,
	0xf9, 0xfb, 0x7d, 0xf6, 0xf1, 0x21, 0xf0, 0x26, 0x94, 0x32, 0x8c, 0x79, 0x6d, 0xca, 0x17, 0xa9,
	0x94, 0x71, 0x52, 0x9b, 0x48, 0x91, 0xb2, 0x48, 0x70, 0xc5, 0x04, 0x8b, 0xaf, 0x93, 0x28, 0xa9,
	0x2d, 0xea, 0x2c, 0x9e, 0x7f, 0x62, 0xf5, 0x5a, 0x34, 0x63, 0x21, 0x1f, 0x8d, 0x59, 0x12, 0x25,
	0xf6, 0x5c, 0xc9, 0x54, 0xe2, 0xa3, 0x5b, 0xb7, 0xbd, 0x72, 0xdb, 0x0f, 0xdc, 0xf6, 0xca, 0x5d,
	0xd9, 0x5f, 0x26, 0xb1, 0x79, 0x54, 0x63, 0x42, 0xc8, 0x94, 0xa5, 0x91, 0x14, 0x4b, 0xd2, 0xc1,
	0x4d, 0x09, 0x2c, 0x57, 0x4e, 0x3e, 0x73, 0xe5, 0xe9, 0x94, 0xca, 0x8f, 0x1c, 0x14, 0x7c, 0x76,
	0xcd, 0x15, 0xe6, 0x60, 0x4e, 0x23, 0xc5, 0x27, 0x69, 0xb4, 0xe0, 0x65, 0xa3, 0x6a, 0x1c, 0xee,
	0x34, 0xda, 0xf6, 0xe3, 0x73, 0xed, 0x7b, 0x54, 0x3b, 0x23, 0xda, 0xee, 0x0a, 0x47, 0xd7, 0x64,
	0xbc, 0x0f, 0x26, 0x53, 0xe1, 0xd5, 0x8c, 0x8b, 0x34, 0x29, 0xe7, 0xaa, 0xc6, 0xa1, 0x49, 0xd7,
	0x85, 0x83, 0xef, 0x06, 0x98, 0x77, 0x36, 0xfc, 0x1c, 0x9e, 0xba, 0x1e, 0x25, 0xce, 0xc0, 0x3b,
	0x23, 0xa3, 0x61, 0xd0, 0xef, 0x11, 0xc7, 0x3b, 0xf6, 0x88, 0x8b, 0x9e, 0xe0, 0x1d, 0x80, 0xd3,
	0xa6, 0x17, 0x0c, 0x9a, 0x5e, 0x40, 0x28, 0x32, 0x70, 0x09, 0xf2, 0x74, 0x18, 0xa0, 0x9c, 0x5e,
	0x38, 0xa7, 0x2e, 0xca, 0x63, 0x13, 0x0a, 0x7e, 0xb3, 0x45, 0x7c, 0xb4, 0x81, 0x01, 0x8a, 0xe4,
	0xbc, 0xd7, 0xed, 0x13, 0x54, 0xd0, 0x3a, 0x09, 0xce, 0x50, 0x51, 0x2f, 0x9a, 0xae, 0x8b, 0x4a,
	0x78, 0x13, 0x36, 0x9c, 0x6e, 0xef, 0x02, 0x6d, 0x6a, 0x28, 0x09, 0x06, 0xf4, 0xa2, 0xd7, 0xf5,
	0x82, 0x01, 0x32, 0xb5, 0xef, 0xac, 0xeb, 0x0f, 0x4f, 0x09, 0x02, 0xbd, 0x6b, 0xd8, 0x27, 0x14,
	0x59, 0xd8, 0x82, 0xd2, 0xbb, 0x2e, 0x3d, 0x71, 0x3d, 0x8a, 0xb6, 0x32, 0x0a, 0x6d, 0xa3, 0x6d,
	0x5d, 0xed, 0x06, 0xad, 0xa1, 0xe7, 0xbb, 0x68, 0x47, 0x83, 0xfa, 0x83, 0x6e, 0xaf, 0xef, 0xb5,
	0x83, 0xa6, 0x8f, 0x76, 0xf1, 0x2e, 0x58, 0x1d, 0xd2, 0xf4, 0x07, 0x1d, 0xa7, 0x43, 0x9c, 0x13,
	0x84, 0xf4, 0xe1, 0xfa, 0x1d, 0xe2, 0xfb, 0x68, 0xaf, 0x72, 0x0e, 0xd6, 0x71, 0x24, 0x42, 0xae,
	0xe6, 0x2a, 0x12, 0x29, 0x7e, 0x06, 0xa5, 0x45, 0x7d, 0x24, 0xd8, 0xec, 0xf6, 0x11, 0x4c, 0x5a,
	0x5c, 0xd4, 0x03, 0x36, 0xe3, 0x99, 0xd0, 0x18, 0x8d, 0x63, 0x39, 0x2e, 0xe7, 0xaa, 0xf9, 0x4c,
	0x68, 0xb4, 0x62, 0x39, 0x5e, 0x0a, 0x99, 0x23, 0xbf, 0x74, 0x34, 0xb4, 0xa3, 0xf2, 0xd5, 0x80,
	0x42, 0x4b, 0x4f, 0x11, 0x7e, 0x01, 0x5b, 0x8a, 0x27, 0xf2, 0x4a, 0x4d, 0xf8, 0xe8, 0x4a, 0xc5,
	0x4b, 0xb2, 0xb5, 0xaa, 0x0d, 0x55, 0x8c, 0x39, 0x58, 0x97, 0xeb, 0x63, 0x64, 0x2f, 0x63, 0x35,
	0x9c, 0xbf, 0x1d, 0x80, 0x7b, 0x37, 0xa2, 0xf7, 0xb9, 0x95, 0x9b, 0x1c, 0x94, 0x5c, 0xae, 0xa2,
	0x05, 0x9f, 0xfe, 0x1e, 0x69, 0xfc, 0x9f, 0x48, 0x5c, 0x81, 0xcd, 0x69, 0x94, 0xa4, 0x4c, 0x4c,
	0x78, 0x76, 0xad, 0x6d, 0x7a, 0xf7, 0x1b, 0x7f, 0x04, 0x88, 0xf5, 0xac, 0x8e, 0x22, 0x71, 0x29,
	0xcb, 0xf9, 0x6a, 0xfe, 0xd0, 0x6a, 0xbc, 0xfd, 0xa7, 0xa9, 0xa7, 0x66, 0x06, 0xf4, 0xc4, 0xa5,
	0xc4, 0x47, 0xb0, 0x37, 0x66, 0x09, 0x1f, 0xfd, 0xd2, 0xfb, 0x8d, 0xac, 0xf7, 0xbb, 0x5a, 0xa0,
	0xeb, 0xfe, 0xb7, 0xbe, 0x18, 0xf0, 0x72, 0x22, 0x67, 0xab, 0xec, 0x3f, 0x47, 0xf6, 0x8c, 0xf7,
	0x1f, 0x96, 0x9b, 0x42, 0x19, 0x33, 0x11, 0xda, 0x52, 0x85, 0xb5, 0x90, 0x8b, 0xec, 0x0f, 0x5e,
	0xbb, 0x95, 0xd8, 0x3c, 0x4a, 0x1e, 0xf3, 0xad, 0x79, 0xfd, 0x40, 0xfa, 0x96, 0xcb, 0xb7, 0x9d,
	0xe6, 0xb8, 0x98, 0xd1, 0x5e, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xd1, 0xf4, 0x9a, 0xb8,
	0x04, 0x00, 0x00,
}
