package service

import "real-time-forum/backend/models"

func (service *AppService) AddReactionLike(reaction *models.Reaction) *models.ErrorJson {
	if err := service.repo.AddReactionLike(reaction); err != nil {
		return err
	}
	return nil
}

func (service *AppService) UpdateReactionLike(reaction *models.Reaction) *models.ErrorJson {
	if err := service.repo.UpdateReactionLike(reaction); err != nil {
		return err
	}
	return nil
}

//

func (service *AppService) React(reaction *models.Reaction, type_entity string) *models.ErrorJson {
	if err := service.repo.CheckEntityID(type_entity, reaction); err != nil {
		return err
	}
	return nil
}

func (service *AppService)  HanldeReaction(reaction *models.Reaction) *models.ErrorJson {
	reaction_existed, err := service.repo.HanldeReaction(reaction)
	if err != nil {
		return &models.ErrorJson{Status: err.Status, Message: err.Message}
	}
	if reaction_existed == nil {
		errJson := service.AddReactionLike(reaction)
		if errJson != nil {
			return errJson
		}
	} else {
		errJson := service.UpdateReactionLike(reaction_existed)
		if errJson != nil {
			return errJson
		}
	}
	return nil
}

// there is a problem
// we need a query to get if there is a reaction for the comment
// each time we need to check the value stored in the database and then
// based on it we chose either to add -1
// like if the first time add 1 and if the second time

func (service *AppService) GetTypeIdByName(type_entity string) int {
	return service.repo.GetTypeIdByName(type_entity)
}
