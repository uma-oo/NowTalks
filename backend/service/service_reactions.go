package service

import "real-time-forum/backend/models"

func (service *AppService) AddReaction(reaction *models.Reaction, reaction_type int) (*models.Reaction, *models.ErrorJson) {
	reaction_created, err := service.repo.AddReaction(reaction, reaction_type)
	if err != nil {
		return nil, err
	}
	return reaction_created, nil
}

func (service *AppService) UpdateReaction(reaction *models.Reaction, reaction_type int) (*models.Reaction, *models.ErrorJson) {
	switch reaction_type {
	case 1:
		reaction_created, err := service.repo.UpdateReactionLike(reaction)
		if err != nil {
			return nil, err
		}
		return reaction_created, nil

	case -1:
		reaction_created, err := service.repo.UpdateReactionDislike(reaction)
		if err != nil {
			return nil, err
		}
		return reaction_created, nil
	}
	return nil, nil
}

//

// func (service *AppService) React(reaction *models.Reaction, type_entity string, reaction_type int) *models.ErrorJson {
// 	if err := service.repo.CheckEntityID(reaction, type_entity); err != nil {
// 		return err
// 	}
// 	return service.HanldeReaction(reaction, reaction_type)
// }

func (service *AppService) HanldeReaction(reaction *models.Reaction, reaction_type int) (*models.Reaction, *models.ErrorJson) {
	reaction_existed, err := service.repo.HanldeReaction(reaction)
	if err != nil {
		return reaction_existed, &models.ErrorJson{Status: err.Status, Message: err.Message}
	}
	if reaction_existed == nil {
		reaction, errJson := service.AddReaction(reaction, reaction_type)
		if errJson != nil {
			return nil, errJson
		}
		return reaction, nil
	} else {
		reaction, errJson := service.UpdateReaction(reaction_existed, reaction_type)
		if errJson != nil {
			return nil, errJson
		}
		return reaction, nil
	}
}

// there is a problem
// we need a query to get if there is a reaction for the comment
// each time we need to check the value stored in the database and then
// based on it we chose either to add -1
// like if the first time add 1 and if the second time

func (service *AppService) GetTypeIdByName(type_entity string) int {
	return service.repo.GetTypeIdByName(type_entity)
}
