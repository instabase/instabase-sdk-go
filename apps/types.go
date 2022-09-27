package apps

// RegionProvenance stores provenance information for regions denoted by
// co-ordinate of top-left and bottom-right pixel of region
type RegionProvenance struct {
	// BoundingPoly denotes the region using top-left and bottom-right co-ordinates
	BoundingPoly struct {
		TopX    float64 `json:"top_x"`
		TopY    float64 `json:"top_y"`
		BottomX float64 `json:"bottom_x"`
		BottomY float64 `json:"bottom_y"`
	} `json:"bounding_poly"`
	// Index represents the page index for the provided region
	Index int64
}

// TextProvenance stores provenance information for text denoted by
// starting and ending position as well as width and height of the text region
type TextProvenance struct {
	Height         float64 `json:"height"`
	Index          float64 `json:"index"`
	OriginalStartX float64 `json:"original_start_x"`
	OriginalStartY float64 `json:"original_start_y"`
	OutputX        float64 `json:"output_x"`
	OutputY        float64 `json:"output_y"`
	Width          float64 `json:"width"`
}

// ProvenanceInfo stores provenance information
type ProvenanceInfo struct {
	ExtractedRegions       []RegionProvenance `json:"extracted_regions"`
	ExtractedTextRegions   []TextProvenance   `json:"extracted_text_regions"`
	InformationRegions     []RegionProvenance `json:"information_regions"`
	InformationTextRegions []TextProvenance   `json:"information_text_regions"`
}

// IndividualKeyValue stores Key, Value and provenance information extracted
// from Instabase
type IndividualKeyValue struct {
	Key        string         `json:"key"`
	Value      string         `json:"value"`
	Provenance ProvenanceInfo `json:"provenance"`
}

// Record represents a list of individual Key Value pairs
type Record struct {
	Results []IndividualKeyValue `json:"results"`
}

// Response stores the error message in case of errors and a list of records
// extracted for the document
type Response struct {
	Error         string   `json:"error,omitempty"`
	Records       []Record `json:"records"`
	APIVersion    string   `json:"api_version"`
	InputFileName string   `json:"input_file_name"`
}
