package graph

import "github.com/fnproject/completer/model"

func internalErrorResult(code model.ErrorDatumType, message string) *model.CompletionResult {
	return &model.CompletionResult{
		Successful: false,
		Datum: &model.Datum{
			Val: &model.Datum_Error{Error: &model.ErrorDatum{Type: code, Message: message}},
		},
	}
}

func emptyDatum() *model.Datum {
	return &model.Datum{Val: &model.Datum_Empty{Empty: &model.EmptyDatum{}}}
}

func blob(contentType string, data []byte) *model.BlobDatum {
	return &model.BlobDatum{ContentType: contentType, DataString: data}
}

func blobDatum(blob *model.BlobDatum) *model.Datum {
	return &model.Datum{
		Val: &model.Datum_Blob{
			Blob: blob,
		},
	}
}

func stageRefDatum(stageID string) *model.Datum {
	return &model.Datum{
		Val: &model.Datum_StageRef{
			StageRef: &model.StageRefDatum{StageRef: stageID},
		},
	}
}

func successfulResult(datum *model.Datum) *model.CompletionResult {
	return &model.CompletionResult{
		Successful: true,
		Datum:      datum,
	}
}

func failedResult(datum *model.Datum) *model.CompletionResult {
	return &model.CompletionResult{
		Successful: false,
		Datum:      datum,
	}
}
