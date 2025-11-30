package domain

type Args struct {
	Size            Size        `json:"size"`
	IterationCount  int         `json:"iteration_count"`
	OutputPath      string      `json:"output_path"`
	Threads         int         `json:"threads"`
	Seed            float64     `json:"seed"`
	Functions       []Function  `json:"functions"`
	AffineParams    AffineParam `json:"affine_params"`
	GammaCorrection bool        `json:"gamma_correction"`
	Gamma           float64     `json:"gamma"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Function struct {
	Name   Transformations `json:"name"`
	Weight float64         `json:"weight"`
}

type AffineParam struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
	D float64 `json:"d"`
	E float64 `json:"e"`
	F float64 `json:"f"`
}
