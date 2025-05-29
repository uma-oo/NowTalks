package service

import "real-time-forum/backend/models"

func (service *AppService) AddReaction(reaction *models.Reaction) *models.ErrorJson {
	if err := service.repo.AddReaction(reaction); err != nil {
		return err
	}
	return nil
}

func (service *AppService) UpdateReaction(reaction *models.Reaction, new_reaction int) *models.ErrorJson {
	if err := service.repo.UpdateReaction(reaction, new_reaction); err != nil {
		return err
	}
	return nil
}
