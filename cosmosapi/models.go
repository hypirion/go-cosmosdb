package cosmosapi

type DataType string

const (
	StringType     = DataType("String")
	NumberType     = DataType("Number")
	PointType      = DataType("Point")
	PolygonType    = DataType("Polygon")
	LineStringType = DataType("LineString")
)

type IndexKind string

const (
	Hash    = IndexKind("Hash")
	Range   = IndexKind("Range")
	Spatial = IndexKind("Spatial")
)

type IndexOrder string

const (
	Ascending  IndexOrder = "ascending"
	Descending IndexOrder = "descending"
)

const MaxPrecision = -1

type Index struct {
	DataType  DataType  `json:"dataType,omitempty"`
	Kind      IndexKind `json:"kind,omitempty"`
	Precision int       `json:"precision,omitempty"`
}

type IncludedPath struct {
	Path    string  `json:"path"`
	Indexes []Index `json:"indexes,omitempty"`
}

type ExcludedPath struct {
	Path string `json:"path"`
}

type CompositeIndex []struct {
	Path  string     `json:"path"`
	Order IndexOrder `json:"order,omitempty"`
}

// Stored Procedure
type Sproc struct {
	Resource
	Body string `json:"body,omitempty"`
}

// User Defined Function
type UDF struct {
	Resource
	Body string `json:"body,omitempty"`
}
