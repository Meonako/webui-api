package api

import (
	"fmt"

	"github.com/goccy/go-json"
)

// Only contains original options (without extension options. Can't posssibly do that.)
type Options struct {
	SamplesSave                        bool     `json:"samples_save,omitempty"`
	SamplesFormat                      string   `json:"samples_format,omitempty"`
	SamplesFilenamePattern             string   `json:"samples_filename_pattern,omitempty"`
	SaveImagesAddNumber                bool     `json:"save_images_add_number,omitempty"`
	GridSave                           bool     `json:"grid_save,omitempty"`
	GridFormat                         string   `json:"grid_format,omitempty"`
	GridExtendedFilename               bool     `json:"grid_extended_filename,omitempty"`
	GridOnlyIfMultiple                 bool     `json:"grid_only_if_multiple,omitempty"`
	GridPreventEmptySpots              bool     `json:"grid_prevent_empty_spots,omitempty"`
	NRows                              int      `json:"n_rows,omitempty"`
	EnablePnginfo                      bool     `json:"enable_pnginfo,omitempty"`
	SaveTxt                            bool     `json:"save_txt,omitempty"`
	SaveImagesBeforeFaceRestoration    bool     `json:"save_images_before_face_restoration,omitempty"`
	SaveImagesBeforeHighresFix         bool     `json:"save_images_before_highres_fix,omitempty"`
	SaveImagesBeforeColorCorrection    bool     `json:"save_images_before_color_correction,omitempty"`
	JpegQuality                        int      `json:"jpeg_quality,omitempty"`
	WebpLossless                       bool     `json:"webp_lossless,omitempty"`
	ExportFor4Chan                     bool     `json:"export_for_4chan,omitempty"`
	ImgDownscaleThreshold              int      `json:"img_downscale_threshold,omitempty"`
	TargetSideLength                   int      `json:"target_side_length,omitempty"`
	ImgMaxSizeMp                       int      `json:"img_max_size_mp,omitempty"`
	UseOriginalNameBatch               bool     `json:"use_original_name_batch,omitempty"`
	UseUpscalerNameAsSuffix            bool     `json:"use_upscaler_name_as_suffix,omitempty"`
	SaveSelectedOnly                   bool     `json:"save_selected_only,omitempty"`
	DoNotAddWatermark                  bool     `json:"do_not_add_watermark,omitempty"`
	TempDir                            string   `json:"temp_dir,omitempty"`
	CleanTempDirAtStart                bool     `json:"clean_temp_dir_at_start,omitempty"`
	OutdirSamples                      string   `json:"outdir_samples,omitempty"`
	OutdirTxt2ImgSamples               string   `json:"outdir_txt2img_samples,omitempty"`
	OutdirImg2ImgSamples               string   `json:"outdir_img2img_samples,omitempty"`
	OutdirExtrasSamples                string   `json:"outdir_extras_samples,omitempty"`
	OutdirGrids                        string   `json:"outdir_grids,omitempty"`
	OutdirTxt2ImgGrids                 string   `json:"outdir_txt2img_grids,omitempty"`
	OutdirImg2ImgGrids                 string   `json:"outdir_img2img_grids,omitempty"`
	OutdirSave                         string   `json:"outdir_save,omitempty"`
	SaveToDirs                         bool     `json:"save_to_dirs,omitempty"`
	GridSaveToDirs                     bool     `json:"grid_save_to_dirs,omitempty"`
	UseSaveToDirsForUI                 bool     `json:"use_save_to_dirs_for_ui,omitempty"`
	DirectoriesFilenamePattern         string   `json:"directories_filename_pattern,omitempty"`
	DirectoriesMaxPromptWords          int      `json:"directories_max_prompt_words,omitempty"`
	ESRGANTile                         int      `json:"ESRGAN_tile,omitempty"`
	ESRGANTileOverlap                  int      `json:"ESRGAN_tile_overlap,omitempty"`
	RealesrganEnabledModels            []string `json:"realesrgan_enabled_models,omitempty"`
	UpscalerForImg2Img                 string   `json:"upscaler_for_img2img,omitempty"`
	FaceRestorationModel               string   `json:"face_restoration_model,omitempty"`
	CodeFormerWeight                   float64  `json:"code_former_weight,omitempty"`
	FaceRestorationUnload              bool     `json:"face_restoration_unload,omitempty"`
	ShowWarnings                       bool     `json:"show_warnings,omitempty"`
	MemmonPollRate                     int      `json:"memmon_poll_rate,omitempty"`
	SamplesLogStdout                   bool     `json:"samples_log_stdout,omitempty"`
	MultipleTqdm                       bool     `json:"multiple_tqdm,omitempty"`
	PrintHypernetExtra                 bool     `json:"print_hypernet_extra,omitempty"`
	UnloadModelsWhenTraining           bool     `json:"unload_models_when_training,omitempty"`
	PinMemory                          bool     `json:"pin_memory,omitempty"`
	SaveOptimizerState                 bool     `json:"save_optimizer_state,omitempty"`
	SaveTrainingSettingsToTxt          bool     `json:"save_training_settings_to_txt,omitempty"`
	DatasetFilenameWordRegex           string   `json:"dataset_filename_word_regex,omitempty"`
	DatasetFilenameJoinString          string   `json:"dataset_filename_join_string,omitempty"`
	TrainingImageRepeatsPerEpoch       int      `json:"training_image_repeats_per_epoch,omitempty"`
	TrainingWriteCsvEvery              int      `json:"training_write_csv_every,omitempty"`
	TrainingXattentionOptimizations    bool     `json:"training_xattention_optimizations,omitempty"`
	TrainingEnableTensorboard          bool     `json:"training_enable_tensorboard,omitempty"`
	TrainingTensorboardSaveImages      bool     `json:"training_tensorboard_save_images,omitempty"`
	TrainingTensorboardFlushEvery      int      `json:"training_tensorboard_flush_every,omitempty"`
	SdModelCheckpoint                  string   `json:"sd_model_checkpoint,omitempty"`
	SdCheckpointCache                  int      `json:"sd_checkpoint_cache,omitempty"`
	SdVaeCheckpointCache               int      `json:"sd_vae_checkpoint_cache,omitempty"`
	SdVae                              string   `json:"sd_vae,omitempty"`
	SdVaeAsDefault                     bool     `json:"sd_vae_as_default,omitempty"`
	InpaintingMaskWeight               int      `json:"inpainting_mask_weight,omitempty"`
	InitialNoiseMultiplier             int      `json:"initial_noise_multiplier,omitempty"`
	Img2ImgColorCorrection             bool     `json:"img2img_color_correction,omitempty"`
	Img2ImgFixSteps                    bool     `json:"img2img_fix_steps,omitempty"`
	Img2ImgBackgroundColor             string   `json:"img2img_background_color,omitempty"`
	EnableQuantization                 bool     `json:"enable_quantization,omitempty"`
	EnableEmphasis                     bool     `json:"enable_emphasis,omitempty"`
	EnableBatchSeeds                   bool     `json:"enable_batch_seeds,omitempty"`
	CommaPaddingBacktrack              int      `json:"comma_padding_backtrack,omitempty"`
	CLIPStopAtLastLayers               int      `json:"CLIP_stop_at_last_layers,omitempty"`
	UpcastAttn                         bool     `json:"upcast_attn,omitempty"`
	UseOldEmphasisImplementation       bool     `json:"use_old_emphasis_implementation,omitempty"`
	UseOldKarrasSchedulerSigmas        bool     `json:"use_old_karras_scheduler_sigmas,omitempty"`
	NoDpmppSdeBatchDeterminism         bool     `json:"no_dpmpp_sde_batch_determinism,omitempty"`
	UseOldHiresFixWidthHeight          bool     `json:"use_old_hires_fix_width_height,omitempty"`
	InterrogateKeepModelsInMemory      bool     `json:"interrogate_keep_models_in_memory,omitempty"`
	InterrogateReturnRanks             bool     `json:"interrogate_return_ranks,omitempty"`
	InterrogateClipNumBeams            int      `json:"interrogate_clip_num_beams,omitempty"`
	InterrogateClipMinLength           int      `json:"interrogate_clip_min_length,omitempty"`
	InterrogateClipMaxLength           int      `json:"interrogate_clip_max_length,omitempty"`
	InterrogateClipDictLimit           int      `json:"interrogate_clip_dict_limit,omitempty"`
	InterrogateClipSkipCategories      []any    `json:"interrogate_clip_skip_categories,omitempty"`
	InterrogateDeepbooruScoreThreshold float64  `json:"interrogate_deepbooru_score_threshold,omitempty"`
	DeepbooruSortAlpha                 bool     `json:"deepbooru_sort_alpha,omitempty"`
	DeepbooruUseSpaces                 bool     `json:"deepbooru_use_spaces,omitempty"`
	DeepbooruEscape                    bool     `json:"deepbooru_escape,omitempty"`
	DeepbooruFilterTags                string   `json:"deepbooru_filter_tags,omitempty"`
	ExtraNetworksDefaultView           string   `json:"extra_networks_default_view,omitempty"`
	ExtraNetworksDefaultMultiplier     int      `json:"extra_networks_default_multiplier,omitempty"`
	ExtraNetworksAddTextSeparator      string   `json:"extra_networks_add_text_separator,omitempty"`
	SdHypernetwork                     string   `json:"sd_hypernetwork,omitempty"`
	ReturnGrid                         bool     `json:"return_grid,omitempty"`
	DoNotShowImages                    bool     `json:"do_not_show_images,omitempty"`
	AddModelHashToInfo                 bool     `json:"add_model_hash_to_info,omitempty"`
	AddModelNameToInfo                 bool     `json:"add_model_name_to_info,omitempty"`
	DisableWeightsAutoSwap             bool     `json:"disable_weights_auto_swap,omitempty"`
	SendSeed                           bool     `json:"send_seed,omitempty"`
	SendSize                           bool     `json:"send_size,omitempty"`
	Font                               string   `json:"font,omitempty"`
	JsModalLightbox                    bool     `json:"js_modal_lightbox,omitempty"`
	JsModalLightboxInitiallyZoomed     bool     `json:"js_modal_lightbox_initially_zoomed,omitempty"`
	ShowProgressInTitle                bool     `json:"show_progress_in_title,omitempty"`
	SamplersInDropdown                 bool     `json:"samplers_in_dropdown,omitempty"`
	DimensionsAndBatchTogether         bool     `json:"dimensions_and_batch_together,omitempty"`
	KeyeditPrecisionAttention          float64  `json:"keyedit_precision_attention,omitempty"`
	KeyeditPrecisionExtra              float64  `json:"keyedit_precision_extra,omitempty"`
	Quicksettings                      string   `json:"quicksettings,omitempty"`
	HiddenTabs                         []any    `json:"hidden_tabs,omitempty"`
	UIReorder                          string   `json:"ui_reorder,omitempty"`
	UIExtraNetworksTabReorder          string   `json:"ui_extra_networks_tab_reorder,omitempty"`
	Localization                       string   `json:"localization,omitempty"`
	ShowProgressbar                    bool     `json:"show_progressbar,omitempty"`
	LivePreviewsEnable                 bool     `json:"live_previews_enable,omitempty"`
	ShowProgressGrid                   bool     `json:"show_progress_grid,omitempty"`
	ShowProgressEveryNSteps            int      `json:"show_progress_every_n_steps,omitempty"`
	ShowProgressType                   string   `json:"show_progress_type,omitempty"`
	LivePreviewContent                 string   `json:"live_preview_content,omitempty"`
	LivePreviewRefreshPeriod           int      `json:"live_preview_refresh_period,omitempty"`
	HideSamplers                       []any    `json:"hide_samplers,omitempty"`
	EtaDdim                            float64  `json:"eta_ddim,omitempty"`
	EtaAncestral                       float64  `json:"eta_ancestral,omitempty"`
	DdimDiscretize                     string   `json:"ddim_discretize,omitempty"`
	SChurn                             float64  `json:"s_churn,omitempty"`
	STmin                              float64  `json:"s_tmin,omitempty"`
	SNoise                             float64  `json:"s_noise,omitempty"`
	EtaNoiseSeedDelta                  int      `json:"eta_noise_seed_delta,omitempty"`
	AlwaysDiscardNextToLastSigma       bool     `json:"always_discard_next_to_last_sigma,omitempty"`
	UniPcVariant                       string   `json:"uni_pc_variant,omitempty"`
	UniPcSkipType                      string   `json:"uni_pc_skip_type,omitempty"`
	UniPcOrder                         int      `json:"uni_pc_order,omitempty"`
	UniPcLowerOrderFinal               bool     `json:"uni_pc_lower_order_final,omitempty"`
	PostprocessingEnableInMainUI       []any    `json:"postprocessing_enable_in_main_ui,omitempty"`
	PostprocessingOperationOrder       []any    `json:"postprocessing_operation_order,omitempty"`
	UpscalingMaxImagesInCache          int      `json:"upscaling_max_images_in_cache,omitempty"`
	DisabledExtensions                 []string `json:"disabled_extensions,omitempty"`
	SdCheckpointHash                   string   `json:"sd_checkpoint_hash,omitempty"`
	SdLora                             string   `json:"sd_lora,omitempty"`
	LoraApplyToOutputs                 bool     `json:"lora_apply_to_outputs,omitempty"`
}

// Get Options.
//
//	NOTE: "NOT" included "extension options".
//	SEE: OptionsWithExtensionOptions() for extension options
func (a *api) Options() (result *Options, err error) {
	resp, erro := a.get(a.Config.Path.Options)
	if erro != nil {
		err = erro
		return
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		err = fmt.Errorf("err: %v\nValue: %v", err, string(resp))
	}
	return
}

// Get Options.
//
//	NOTE: Return as a map. SO you need to know the key(s) and value(s) type.
func (a *api) OptionsWithExtensionOptions() (result map[string]any, err error) {
	resp, erro := a.get(a.Config.Path.Options)
	if erro != nil {
		err = erro
		return
	}

	err = json.Unmarshal(resp, &result)
	return
}

// Set original options.
//
//	SEE: SetOptionsWithExtensionOptions() for extension options settings.
func (a *api) SetOptions(params *Options) error {
	payload, err := json.Marshal(params)
	if err != nil {
		return err
	}

	_, err = a.post(a.Config.Path.Options, payload)
	return err
}

// Shorthand for map[string]any
type Opt map[string]any

// Receive api.Opt (map[string]any) as argument.
func (a *api) SetOptionsWithExtensionOptions(params Opt) error {
	payload, err := json.Marshal(params)
	if err != nil {
		return err
	}

	_, err = a.post(a.Config.Path.Options, payload)
	return err
}
