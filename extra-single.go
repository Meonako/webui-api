package api

import "github.com/goccy/go-json"

type ExtraSingleImage struct {
	// Resize Mode: Scale By || Scale To. See: webup-api/extra package for helper.
	//
	//  Default: 0 (Scale By)
	ResizeMode int `json:"resize_mode,omitempty"`

	// Don't even know what this is.
	//
	//  Default: true
	ShowExtrasResults bool `json:"show_extras_results,omitempty"`

	// GFPGAN Face restoration. Value must be between 0.0 - 1.0
	//
	//  Default: 0.0
	GfpganVisibility float64 `json:"gfpgan_visibility,omitempty"`

	// CodeFormer Face restoration. Value must be between 0.0 - 1.0
	//
	//  Default: 0.0
	CodeformerVisibility float64 `json:"codeformer_visibility,omitempty"`

	// CodeFormer Face restoration weight. 0 = Maximum Effect, 1 = Minimum Effect.
	//
	//  Default: 0.0
	CodeformerWeight float64 `json:"codeformer_weight,omitempty"`

	// Multiplier to width and height of the original image.
	//
	//  NOTE: Will only work if ResizeMode is 0
	//  Default: 2
	UpscalingResize int `json:"upscaling_resize,omitempty"`

	// Width of the result image.
	//
	//  NOTE: Will only work if ResizeMode is 1
	//  Default: 512
	UpscalingResizeW int `json:"upscaling_resize_w,omitempty"`

	// Height of the result image.
	//
	//  NOTE: Will only work if ResizeMode is 1
	//  Default: 512
	UpscalingResizeH int `json:"upscaling_resize_h,omitempty"`

	// Crop Image if the aspect ratio of original image and result image doesn't match.
	//
	//  NOTE: Will only work if ResizeMode is 1
	//  Default: true
	UpscalingCrop bool `json:"upscaling_crop,omitempty"`

	// First Upscaler Model. See: webui-api/extra package for helper.
	//
	//  Default: "None"
	Upscaler1 string `json:"upscaler_1,omitempty"`

	// Second Upscaler Model. See: webui-api/extra package for helper.
	//
	//  Default: "None"
	Upscaler2 string `json:"upscaler_2,omitempty"`

	// Second Upscaler Model Visibility. See: webui-api/extra package for helper.
	//
	//  Default: 0.0
	ExtrasUpscaler2Visibility float64 `json:"extras_upscaler_2_visibility,omitempty"`

	// Upscale first then do face restoration.
	//
	//  Default: false
	UpscaleFirst bool `json:"upscale_first,omitempty"`

	// Base64-encoded image to be upscale.
	//
	//  Default: ""
	Image string `json:"image,omitempty"`

	// If true, Will Decode Images after received response from API
	DecodeAfterResult bool `json:"-"`
}

func (a *api) ExtraSingleImage(params ExtraSingleImage) (*extraSingleImageRespond, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return &extraSingleImageRespond{}, err
	}

	data, err := a.post(a.Config.Path.ExtraSingle, payload)
	if err != nil {
		return &extraSingleImageRespond{}, err
	}

	apiResp := extraSingleImageRespond{}
	err = json.Unmarshal(data, &apiResp)
	if err != nil {
		return &extraSingleImageRespond{}, err
	}

	if params.DecodeAfterResult {
		_, err := apiResp.DecodeImage()
		if err != nil {
			return &extraSingleImageRespond{}, err
		}
	}

	return &apiResp, nil
}
