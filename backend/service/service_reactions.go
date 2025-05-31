package service

import "real-time-forum/backend/models"

func (service *AppService) AddReaction(reaction *models.Reaction, reaction_type int) *models.ErrorJson {
	if err := service.repo.AddReaction(reaction, reaction_type); err != nil {
		return err
	}
	return nil
}

func (service *AppService) UpdateReaction(reaction *models.Reaction, reaction_type int) *models.ErrorJson {
	switch reaction_type {
	case 1:
		if err := service.repo.UpdateReactionLike(reaction); err != nil {
			return err
		}
	case -1:
		if err := service.repo.UpdateReactionDislike(reaction); err != nil {
			return err
		}
	}
	return nil
}

//

func (service *AppService) React(reaction *models.Reaction, type_entity string, reaction_type int) *models.ErrorJson {
	if err := service.repo.CheckEntityID(reaction, type_entity); err != nil {
		return err
	}
	return service.HanldeReaction(reaction, reaction_type)
}

func (service *AppService) HanldeReaction(reaction *models.Reaction, reaction_type int) *models.ErrorJson {
	reaction_existed, err := service.repo.HanldeReaction(reaction)
	if err != nil {
		return &models.ErrorJson{Status: err.Status, Message: err.Message}
	}
	if reaction_existed == nil {
		errJson := service.AddReaction(reaction, reaction_type)
		if errJson != nil {
			return errJson
		}
	} else {
		errJson := service.UpdateReaction(reaction_existed, reaction_type)
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
