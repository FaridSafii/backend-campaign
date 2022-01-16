package campaign

import "gorm.io/gorm"

type Repository interface {
	//[]Capaign untuk menampung data lebih banyak
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	//jangan lupakan &
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	//Preload digunakan untuk relasi
	//CampaignImages nama di sruct Campaign
	//campaign_images =>nama tabel , is_primary =>nama fieldnya
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
