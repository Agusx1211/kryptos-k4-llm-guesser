package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GenerationSettings struct {
	NCtx                int      `json:"n_ctx"`
	NPredict            int      `json:"n_predict"`
	Model               string   `json:"model"`
	Seed                uint32   `json:"seed"`
	Temperature         float32  `json:"temperature"`
	DynatempRange       float32  `json:"dynatemp_range"`
	DynatempExponent    float32  `json:"dynatemp_exponent"`
	TopK                float32  `json:"top_k"`
	TopP                float32  `json:"top_p"`
	MinP                float32  `json:"min_p"`
	TFSZ                float32  `json:"tfs_z"`
	TypicalP            float32  `json:"typical_p"`
	RepeatLastN         float32  `json:"repeat_last_n"`
	RepeatPenalty       float32  `json:"repeat_penalty"`
	PresencePenalty     float32  `json:"presence_penalty"`
	FrequencyPenalty    float32  `json:"frequency_penalty"`
	PenaltyPromptTokens []string `json:"penalty_prompt_tokens"`
	UsePenaltyPrompt    bool     `json:"use_penalty_prompt_tokens"`
	Mirostat            float32  `json:"mirostat"`
	MirostatTau         float32  `json:"mirostat_tau"`
	MirostatEta         float32  `json:"mirostat_eta"`
	PenalizeNL          bool     `json:"penalize_nl"`
	Stop                []string `json:"stop"`
	NKeep               int      `json:"n_keep"`
	NDiscard            int      `json:"n_discard"`
	IgnoreEOS           bool     `json:"ignore_eos"`
	Stream              bool     `json:"stream"`
	LogitBias           []string `json:"logit_bias"`
	NProbs              int      `json:"n_probs"`
	MinKeep             int      `json:"min_keep"`
	Grammar             string   `json:"grammar"`
	Samplers            []string `json:"samplers"`
}

type CompletionProbabilities struct {
	Content string `json:"content"`
	Probs   []struct {
		TokenStr string  `json:"tok_str"`
		Prob     float64 `json:"prob"`
	} `json:"probs"`
}

type CompletionResponse struct {
	Content            string                    `json:"content"`
	IDSlot             int                       `json:"id_slot"`
	Stop               bool                      `json:"stop"`
	Model              string                    `json:"model"`
	TokensPredicted    int                       `json:"tokens_predicted"`
	TokensEvaluated    int                       `json:"tokens_evaluated"`
	GenerationSettings GenerationSettings        `json:"generation_settings"`
	Prompt             string                    `json:"prompt"`
	Truncated          bool                      `json:"truncated"`
	StoppedEOS         bool                      `json:"stopped_eos"`
	StoppedWord        bool                      `json:"stopped_word"`
	StoppedLimit       bool                      `json:"stopped_limit"`
	StoppingWord       string                    `json:"stopping_word"`
	TokensCached       int                       `json:"tokens_cached"`
	Timings            map[string]float64        `json:"timings"`
	CompletionProbs    []CompletionProbabilities `json:"completion_probabilities"`
}

type CompletionRequest struct {
	Prefix   string `json:"input_prefix"`
	Suffix   string `json:"input_suffix"`
	Prompt   string `json:"prompt"`
	NPredict int    `json:"n_predict"`
	NProbs   int    `json:"n_probs"`
}

type Completer struct {
	URL     string
	limiter chan struct{}
}

func NewCompleter(url string, jobs int) *Completer {
	return &Completer{
		URL:     url,
		limiter: make(chan struct{}, jobs),
	}
}

func (c *Completer) MakeCompletionRequest(prompt string, suffix string, nPredict, nProbs int) (*CompletionResponse, error) {
	c.limiter <- struct{}{}
	defer func() {
		<-c.limiter
	}()

	fullUrl := c.URL + "completion"
	reqBody := CompletionRequest{
		Prompt:   prompt,
		NPredict: nPredict,
		NProbs:   nProbs,
	}

	if suffix != "" {
		reqBody.Prefix = prompt
		reqBody.Suffix = suffix
		reqBody.Prompt = ""
		fullUrl = c.URL + "infill"
	}

	reqData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	resp, err := http.Post(fullUrl, "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var completionResp CompletionResponse
	if err := json.Unmarshal(respBody, &completionResp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	return &completionResp, nil
}
