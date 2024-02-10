package computervision

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"io"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.1/computervision"

// AdultInfo an object describing whether the image contains adult-oriented content and/or is racy.
type AdultInfo struct {
	// IsAdultContent - A value indicating if the image contains adult-oriented content.
	IsAdultContent *bool `json:"isAdultContent,omitempty"`
	// IsRacyContent - A value indicating if the image is racy.
	IsRacyContent *bool `json:"isRacyContent,omitempty"`
	// IsGoryContent - A value indicating if the image is gory.
	IsGoryContent *bool `json:"isGoryContent,omitempty"`
	// AdultScore - Score from 0 to 1 that indicates how much the content is considered adult-oriented within the image.
	AdultScore *float64 `json:"adultScore,omitempty"`
	// RacyScore - Score from 0 to 1 that indicates how suggestive is the image.
	RacyScore *float64 `json:"racyScore,omitempty"`
	// GoreScore - Score from 0 to 1 that indicates how gory is the image.
	GoreScore *float64 `json:"goreScore,omitempty"`
}

// AreaOfInterestResult result of AreaOfInterest operation.
type AreaOfInterestResult struct {
	autorest.Response `json:"-"`
	// AreaOfInterest - READ-ONLY; A bounding box for an area of interest inside an image.
	AreaOfInterest *BoundingRect `json:"areaOfInterest,omitempty"`
	// RequestID - Id of the REST API request.
	RequestID *string        `json:"requestId,omitempty"`
	Metadata  *ImageMetadata `json:"metadata,omitempty"`
}

// MarshalJSON is the custom marshaler for AreaOfInterestResult.
func (aoir AreaOfInterestResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aoir.RequestID != nil {
		objectMap["requestId"] = aoir.RequestID
	}
	if aoir.Metadata != nil {
		objectMap["metadata"] = aoir.Metadata
	}
	return json.Marshal(objectMap)
}

// BoundingRect a bounding box for an area inside an image.
type BoundingRect struct {
	// X - X-coordinate of the top left point of the area, in pixels.
	X *int32 `json:"x,omitempty"`
	// Y - Y-coordinate of the top left point of the area, in pixels.
	Y *int32 `json:"y,omitempty"`
	// W - Width measured from the top-left point of the area, in pixels.
	W *int32 `json:"w,omitempty"`
	// H - Height measured from the top-left point of the area, in pixels.
	H *int32 `json:"h,omitempty"`
}

// Category an object describing identified category.
type Category struct {
	// Name - Name of the category.
	Name *string `json:"name,omitempty"`
	// Score - Scoring of the category.
	Score *float64 `json:"score,omitempty"`
	// Detail - Details of the identified category.
	Detail *CategoryDetail `json:"detail,omitempty"`
}

// CategoryDetail an object describing additional category details.
type CategoryDetail struct {
	// Celebrities - An array of celebrities if any identified.
	Celebrities *[]CelebritiesModel `json:"celebrities,omitempty"`
	// Landmarks - An array of landmarks if any identified.
	Landmarks *[]LandmarksModel `json:"landmarks,omitempty"`
}

// CelebritiesModel an object describing possible celebrity identification.
type CelebritiesModel struct {
	// Name - Name of the celebrity.
	Name *string `json:"name,omitempty"`
	// Confidence - Confidence level for the celebrity recognition as a value ranging from 0 to 1.
	Confidence *float64 `json:"confidence,omitempty"`
	// FaceRectangle - Location of the identified face in the image.
	FaceRectangle *FaceRectangle `json:"faceRectangle,omitempty"`
}

// CelebrityResults result of domain-specific classifications for the domain of celebrities.
type CelebrityResults struct {
	// Celebrities - List of celebrities recognized in the image.
	Celebrities *[]CelebritiesModel `json:"celebrities,omitempty"`
	// RequestID - Id of the REST API request.
	RequestID *string        `json:"requestId,omitempty"`
	Metadata  *ImageMetadata `json:"metadata,omitempty"`
}

// ColorInfo an object providing additional metadata describing color attributes.
type ColorInfo struct {
	// DominantColorForeground - Possible dominant foreground color.
	DominantColorForeground *string `json:"dominantColorForeground,omitempty"`
	// DominantColorBackground - Possible dominant background color.
	DominantColorBackground *string `json:"dominantColorBackground,omitempty"`
	// DominantColors - An array of possible dominant colors.
	DominantColors *[]string `json:"dominantColors,omitempty"`
	// AccentColor - Possible accent color.
	AccentColor *string `json:"accentColor,omitempty"`
	// IsBWImg - A value indicating if the image is black and white.
	IsBWImg *bool `json:"isBWImg,omitempty"`
}

// DetectedBrand a brand detected in an image.
type DetectedBrand struct {
	// Name - READ-ONLY; Label for the brand.
	Name *string `json:"name,omitempty"`
	// Confidence - READ-ONLY; Confidence score of having observed the brand in the image, as a value ranging from 0 to 1.
	Confidence *float64 `json:"confidence,omitempty"`
	// Rectangle - READ-ONLY; Approximate location of the detected brand.
	Rectangle *BoundingRect `json:"rectangle,omitempty"`
}

// MarshalJSON is the custom marshaler for DetectedBrand.
func (db DetectedBrand) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// DetectedObject an object detected in an image.
type DetectedObject struct {
	// Rectangle - READ-ONLY; Approximate location of the detected object.
	Rectangle *BoundingRect `json:"rectangle,omitempty"`
	// Object - Label for the object.
	Object *string `json:"object,omitempty"`
	// Confidence - Confidence score of having observed the object in the image, as a value ranging from 0 to 1.
	Confidence *float64 `json:"confidence,omitempty"`
	// Parent - The parent object, from a taxonomy perspective.
	// The parent object is a more generic form of this object.  For example, a 'bulldog' would have a parent of 'dog'.
	Parent *ObjectHierarchy `json:"parent,omitempty"`
}

// MarshalJSON is the custom marshaler for DetectedObject.
func (do DetectedObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if do.Object != nil {
		objectMap["object"] = do.Object
	}
	if do.Confidence != nil {
		objectMap["confidence"] = do.Confidence
	}
	if do.Parent != nil {
		objectMap["parent"] = do.Parent
	}
	return json.Marshal(objectMap)
}

// DetectResult result of a DetectImage call.
type DetectResult struct {
	autorest.Response `json:"-"`
	// Objects - READ-ONLY; An array of detected objects.
	Objects *[]DetectedObject `json:"objects,omitempty"`
	// RequestID - Id of the REST API request.
	RequestID *string        `json:"requestId,omitempty"`
	Metadata  *ImageMetadata `json:"metadata,omitempty"`
}

// MarshalJSON is the custom marshaler for DetectResult.
func (dr DetectResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dr.RequestID != nil {
		objectMap["requestId"] = dr.RequestID
	}
	if dr.Metadata != nil {
		objectMap["metadata"] = dr.Metadata
	}
	return json.Marshal(objectMap)
}

// DomainModelResults result of image analysis using a specific domain model including additional metadata.
type DomainModelResults struct {
	autorest.Response `json:"-"`
	// Result - Model-specific response.
	Result interface{} `json:"result,omitempty"`
	// RequestID - Id of the REST API request.
	RequestID *string        `json:"requestId,omitempty"`
	Metadata  *ImageMetadata `json:"metadata,omitempty"`
}

// Error details about the API request error.
type Error struct {
	// Code - The error code.
	Code interface{} `json:"code,omitempty"`
	// Message - A message explaining the error reported by the service.
	Message *string `json:"message,omitempty"`
	// RequestID - A unique request identifier.
	RequestID *string `json:"requestId,omitempty"`
}

// FaceDescription an object describing a face identified in the image.
type FaceDescription struct {
	// Age - Possible age of the face.
	Age *int32 `json:"age,omitempty"`
	// Gender - Possible gender of the face. Possible values include: 'Male', 'Female'
	Gender Gender `json:"gender,omitempty"`
	// FaceRectangle - Rectangle in the image containing the identified face.
	FaceRectangle *FaceRectangle `json:"faceRectangle,omitempty"`
}

// FaceRectangle an object describing face rectangle.
type FaceRectangle struct {
	// Left - X-coordinate of the top left point of the face, in pixels.
	Left *int32 `json:"left,omitempty"`
	// Top - Y-coordinate of the top left point of the face, in pixels.
	Top *int32 `json:"top,omitempty"`
	// Width - Width measured from the top-left point of the face, in pixels.
	Width *int32 `json:"width,omitempty"`
	// Height - Height measured from the top-left point of the face, in pixels.
	Height *int32 `json:"height,omitempty"`
}

// ImageAnalysis result of AnalyzeImage operation.
type ImageAnalysis struct {
	autorest.Response `json:"-"`
	// Categories - An array indicating identified categories.
	Categories *[]Category `json:"categories,omitempty"`
	// Adult - An object describing whether the image contains adult-oriented content and/or is racy.
	Adult *AdultInfo `json:"adult,omitempty"`
	// Color - An object providing additional metadata describing color attributes.
	Color *ColorInfo `json:"color,omitempty"`
	// ImageType - An object providing possible image types and matching confidence levels.
	ImageType *ImageType `json:"imageType,omitempty"`
	// Tags - A list of tags with confidence level.
	Tags *[]ImageTag `json:"tags,omitempty"`
	// Description - A collection of content tags, along with a list of captions sorted by confidence level, and image metadata.
	Description *ImageDescriptionDetails `json:"description,omitempty"`
	// Faces - An array of possible faces within the image.
	Faces *[]FaceDescription `json:"faces,omitempty"`
	// Objects - Array of objects describing what was detected in the image.
	Objects *[]DetectedObject `json:"objects,omitempty"`
	// Brands - Array of brands detected in the image.
	Brands *[]DetectedBrand `json:"brands,omitempty"`
	// RequestID - Id of the REST API request.
	RequestID *string        `json:"requestId,omitempty"`
	Metadata  *ImageMetadata `json:"metadata,omitempty"`
}

// ImageCaption an image caption, i.e. a brief description of what the image depicts.
type ImageCaption struct {
	// Text - The text of the caption.
	Text *string `json:"text,omitempty"`
	// Confidence - The level of confidence the service has in the caption.
	Confidence *float64 `json:"confidence,omitempty"`
}

// ImageDescription a collection of content tags, along with a list of captions sorted by confidence level,
// and image metadata.
type ImageDescription struct {
	autorest.Response `json:"-"`
	// ImageDescriptionDetails - A collection of content tags, along with a list of captions sorted by confidence level, and image metadata.
	*ImageDescriptionDetails `json:"description,omitempty"`
	// RequestID - Id of the REST API request.
	RequestID *string        `json:"requestId,omitempty"`
	Metadata  *ImageMetadata `json:"metadata,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageDescription.
func (ID ImageDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ID.ImageDescriptionDetails != nil {
		objectMap["description"] = ID.ImageDescriptionDetails
	}
	if ID.RequestID != nil {
		objectMap["requestId"] = ID.RequestID
	}
	if ID.Metadata != nil {
		objectMap["metadata"] = ID.Metadata
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ImageDescription struct.
func (ID *ImageDescription) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "description":
			if v != nil {
				var imageDescriptionDetails ImageDescriptionDetails
				err = json.Unmarshal(*v, &imageDescriptionDetails)
				if err != nil {
					return err
				}
				ID.ImageDescriptionDetails = &imageDescriptionDetails
			}
		case "requestId":
			if v != nil {
				var requestID string
				err = json.Unmarshal(*v, &requestID)
				if err != nil {
					return err
				}
				ID.RequestID = &requestID
			}
		case "metadata":
			if v != nil {
				var metadata ImageMetadata
				err = json.Unmarshal(*v, &metadata)
				if err != nil {
					return err
				}
				ID.Metadata = &metadata
			}
		}
	}

	return nil
}

// ImageDescriptionDetails a collection of content tags, along with a list of captions sorted by confidence
// level, and image metadata.
type ImageDescriptionDetails struct {
	// Tags - A collection of image tags.
	Tags *[]string `json:"tags,omitempty"`
	// Captions - A list of captions, sorted by confidence level.
	Captions *[]ImageCaption `json:"captions,omitempty"`
}

// ImageMetadata image metadata.
type ImageMetadata struct {
	// Width - Image width, in pixels.
	Width *int32 `json:"width,omitempty"`
	// Height - Image height, in pixels.
	Height *int32 `json:"height,omitempty"`
	// Format - Image format.
	Format *string `json:"format,omitempty"`
}

// ImageTag an entity observation in the image, along with the confidence score.
type ImageTag struct {
	// Name - Name of the entity.
	Name *string `json:"name,omitempty"`
	// Confidence - The level of confidence that the entity was observed.
	Confidence *float64 `json:"confidence,omitempty"`
	// Hint - Optional hint/details for this tag.
	Hint *string `json:"hint,omitempty"`
}

// ImageType an object providing possible image types and matching confidence levels.
type ImageType struct {
	// ClipArtType - Confidence level that the image is a clip art.
	ClipArtType *int32 `json:"clipArtType,omitempty"`
	// LineDrawingType - Confidence level that the image is a line drawing.
	LineDrawingType *int32 `json:"lineDrawingType,omitempty"`
}

// ImageURL ...
type ImageURL struct {
	// URL - Publicly reachable URL of an image.
	URL *string `json:"url,omitempty"`
}

// LandmarkResults result of domain-specific classifications for the domain of landmarks.
type LandmarkResults struct {
	// Landmarks - List of landmarks recognized in the image.
	Landmarks *[]LandmarksModel `json:"landmarks,omitempty"`
	// RequestID - Id of the REST API request.
	RequestID *string        `json:"requestId,omitempty"`
	Metadata  *ImageMetadata `json:"metadata,omitempty"`
}

// LandmarksModel a landmark recognized in the image.
type LandmarksModel struct {
	// Name - Name of the landmark.
	Name *string `json:"name,omitempty"`
	// Confidence - Confidence level for the landmark recognition as a value ranging from 0 to 1.
	Confidence *float64 `json:"confidence,omitempty"`
}

// Line an object representing a recognized text line.
type Line struct {
	// BoundingBox - Bounding box of a recognized line.
	BoundingBox *[]float64 `json:"boundingBox,omitempty"`
	// Text - The text content of the line.
	Text *string `json:"text,omitempty"`
	// Words - List of words in the text line.
	Words *[]Word `json:"words,omitempty"`
}

// ListModelsResult result of the List Domain Models operation.
type ListModelsResult struct {
	autorest.Response `json:"-"`
	// ModelsProperty - READ-ONLY; An array of supported models.
	ModelsProperty *[]ModelDescription `json:"models,omitempty"`
}

// MarshalJSON is the custom marshaler for ListModelsResult.
func (lmr ListModelsResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ModelDescription an object describing supported model by name and categories.
type ModelDescription struct {
	// Name - The name of the model.
	Name *string `json:"name,omitempty"`
	// Categories - Categories of the model.
	Categories *[]string `json:"categories,omitempty"`
}

// ObjectHierarchy an object detected inside an image.
type ObjectHierarchy struct {
	// Object - Label for the object.
	Object *string `json:"object,omitempty"`
	// Confidence - Confidence score of having observed the object in the image, as a value ranging from 0 to 1.
	Confidence *float64 `json:"confidence,omitempty"`
	// Parent - The parent object, from a taxonomy perspective.
	// The parent object is a more generic form of this object.  For example, a 'bulldog' would have a parent of 'dog'.
	Parent *ObjectHierarchy `json:"parent,omitempty"`
}

// OcrLine an object describing a single recognized line of text.
type OcrLine struct {
	// BoundingBox - Bounding box of a recognized line. The four integers represent the x-coordinate of the left edge, the y-coordinate of the top edge, width, and height of the bounding box, in the coordinate system of the input image, after it has been rotated around its center according to the detected text angle (see textAngle property), with the origin at the top-left corner, and the y-axis pointing down.
	BoundingBox *string `json:"boundingBox,omitempty"`
	// Words - An array of objects, where each object represents a recognized word.
	Words *[]OcrWord `json:"words,omitempty"`
}

// OcrRegion a region consists of multiple lines (e.g. a column of text in a multi-column document).
type OcrRegion struct {
	// BoundingBox - Bounding box of a recognized region. The four integers represent the x-coordinate of the left edge, the y-coordinate of the top edge, width, and height of the bounding box, in the coordinate system of the input image, after it has been rotated around its center according to the detected text angle (see textAngle property), with the origin at the top-left corner, and the y-axis pointing down.
	BoundingBox *string `json:"boundingBox,omitempty"`
	// Lines - An array of recognized lines of text.
	Lines *[]OcrLine `json:"lines,omitempty"`
}

// OcrResult ...
type OcrResult struct {
	autorest.Response `json:"-"`
	// Language - The BCP-47 language code of the text in the image.
	Language *string `json:"language,omitempty"`
	// TextAngle - The angle, in radians, of the detected text with respect to the closest horizontal or vertical direction. After rotating the input image clockwise by this angle, the recognized text lines become horizontal or vertical. In combination with the orientation property it can be used to overlay recognition results correctly on the original image, by rotating either the original image or recognition results by a suitable angle around the center of the original image. If the angle cannot be confidently detected, this property is not present. If the image contains text at different angles, only part of the text will be recognized correctly.
	TextAngle *float64 `json:"textAngle,omitempty"`
	// Orientation - Orientation of the text recognized in the image, if requested. The value (up, down, left, or right) refers to the direction that the top of the recognized text is facing, after the image has been rotated around its center according to the detected text angle (see textAngle property).
	// If detection of the orientation was not requested, or no text is detected, the value is 'NotDetected'.
	Orientation *string `json:"orientation,omitempty"`
	// Regions - An array of objects, where each object represents a region of recognized text.
	Regions *[]OcrRegion `json:"regions,omitempty"`
}

// OcrWord information on a recognized word.
type OcrWord struct {
	// BoundingBox - Bounding box of a recognized word. The four integers represent the x-coordinate of the left edge, the y-coordinate of the top edge, width, and height of the bounding box, in the coordinate system of the input image, after it has been rotated around its center according to the detected text angle (see textAngle property), with the origin at the top-left corner, and the y-axis pointing down.
	BoundingBox *string `json:"boundingBox,omitempty"`
	// Text - String value of a recognized word.
	Text *string `json:"text,omitempty"`
}

// ReadCloser ...
type ReadCloser struct {
	autorest.Response `json:"-"`
	Value             *io.ReadCloser `json:"value,omitempty"`
}

// ReadOperationResult OCR result of the read operation.
type ReadOperationResult struct {
	autorest.Response `json:"-"`
	// Status - Status of the read operation. Possible values include: 'NotStarted', 'Running', 'Failed', 'Succeeded'
	Status TextOperationStatusCodes `json:"status,omitempty"`
	// RecognitionResults - An array of text recognition result of the read operation.
	RecognitionResults *[]TextRecognitionResult `json:"recognitionResults,omitempty"`
}

// TagResult the results of a image tag operation, including any tags and image metadata.
type TagResult struct {
	autorest.Response `json:"-"`
	// Tags - A list of tags with confidence level.
	Tags *[]ImageTag `json:"tags,omitempty"`
	// RequestID - Id of the REST API request.
	RequestID *string        `json:"requestId,omitempty"`
	Metadata  *ImageMetadata `json:"metadata,omitempty"`
}

// TextOperationResult result of recognition text operation.
type TextOperationResult struct {
	autorest.Response `json:"-"`
	// Status - Status of the text operation. Possible values include: 'NotStarted', 'Running', 'Failed', 'Succeeded'
	Status TextOperationStatusCodes `json:"status,omitempty"`
	// RecognitionResult - Text recognition result of the text operation.
	RecognitionResult *TextRecognitionResult `json:"recognitionResult,omitempty"`
}

// TextRecognitionResult an object representing a recognized text region
type TextRecognitionResult struct {
	// Page - The 1-based page number of the recognition result.
	Page *int32 `json:"page,omitempty"`
	// ClockwiseOrientation - The orientation of the image in degrees in the clockwise direction. Range between [0, 360).
	ClockwiseOrientation *float64 `json:"clockwiseOrientation,omitempty"`
	// Width - The width of the image in pixels or the PDF in inches.
	Width *float64 `json:"width,omitempty"`
	// Height - The height of the image in pixels or the PDF in inches.
	Height *float64 `json:"height,omitempty"`
	// Unit - The unit used in the Width, Height and BoundingBox. For images, the unit is 'pixel'. For PDF, the unit is 'inch'. Possible values include: 'Pixel', 'Inch'
	Unit TextRecognitionResultDimensionUnit `json:"unit,omitempty"`
	// Lines - A list of recognized text lines.
	Lines *[]Line `json:"lines,omitempty"`
}

// Word an object representing a recognized word.
type Word struct {
	// BoundingBox - Bounding box of a recognized word.
	BoundingBox *[]float64 `json:"boundingBox,omitempty"`
	// Text - The text content of the word.
	Text *string `json:"text,omitempty"`
	// Confidence - Qualitative confidence measure. Possible values include: 'High', 'Low'
	Confidence TextRecognitionResultConfidenceClass `json:"confidence,omitempty"`
}
