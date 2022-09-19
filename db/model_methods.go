package db

import "quiz/dto"

func (q Quiz) GetDTO() dto.QuizResponseDTO {
	return dto.QuizResponseDTO{
		Name:          q.Name,
		Completed:     q.Completed,
		ID:            q.ID,
		QuestionCount: q.QuestionCount,
		ThemeCount:    q.ThemeCount,
		SuperRound:    q.SuperRound,
	}
}

func (t Theme) GetDTO() dto.ThemeResponseDTO {
	return dto.ThemeResponseDTO{
		Name:  t.Name,
		ID:    t.ID,
		Round: int(t.Round),
	}
}

func (q Question) GetDTO() dto.QuestionResponseDTO {
	return dto.QuestionResponseDTO{
		TextContent:       q.TextContent,
		ID:                q.ID,
		Value:             q.Value,
		MultimediaContent: q.MultimediaContent,
		MultimediaType:    q.MultimediaType,
	}
}
