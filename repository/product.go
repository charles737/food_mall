package repository

import (
	"fmt"
	"food_mall/model"
	"food_mall/query"
	"food_mall/utils"
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

type ProductRepoInterface interface {
	List(req *query.ListQuery) (Products []*model.Product, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(Product model.Product) (*model.Product, error)
	Exist(Product model.Product) *model.Product
	ExistByProductID(id string) *model.Product
	Add(Product model.Product) (*model.Product, error)
	Edit(Product model.Product) (bool, error)
	Delete(p model.Product) (bool, error)
}

func (repo *ProductRepository) List(req *query.ListQuery) (products []*model.Product, err error) {
	fmt.Println(req)

	// 分页
	limit, offset := utils.Page(req.PageSize, req.Page)
	if err := repo.DB.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil

}

func (repo *ProductRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var products []*model.Product
	if err := repo.DB.Find(products).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *ProductRepository) Get(product model.Product) (*model.Product, error) {
	if err := repo.DB.Where(&product).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) Exist(product model.Product) *model.Product {
	if product.ProductName != "" {
		var temp model.Product
		repo.DB.Where("product_name=?", product.ProductName).First(&temp)
		return &temp
	}
	return nil
}

func (repo *ProductRepository) ExistByProductID(id string) *model.Product {
	var p model.Product
	repo.DB.Where("product_id=?", id).First(&p)
	return &p
}

func (repo *ProductRepository) Add(product model.Product) (*model.Product, error) {
	exist := repo.Exist(product)
	if exist != nil && exist.ProductName != "" {
		return &product, fmt.Errorf("商品已存在")
	}
	if err := repo.DB.Create(&product).Error; err != nil {
		return nil, fmt.Errorf("商品添加失败")
	}
	return &product, nil
}

func (repo *ProductRepository) Edit(product model.Product) (bool, error) {
	if product.ProductId == "" {
		return false, fmt.Errorf("请输入要编辑的商品ID")
	}
	err := repo.DB.Model(&product).Where("product_id=?", product.ProductId).Updates(map[string]interface{}{
		"product_name": product.ProductName,
		"product_intro": product.ProductIntro,
		"category_id": product.CategoryId,
		"product_cover_img": product.ProductCoverImg,
		"product_banner": product.ProductBanner,
		"original_price": product.OriginalPrice,
		"selling_price": product.SellingPrice,
		"stock_num": product.StockNum,
		"tag": product.Tag,
		"sell_status": product.SellStatus,
		"product_detail_content": product.ProductDetailContent,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *ProductRepository) Delete(p model.Product) (bool, error) {
	err := repo.DB.Model(&p).Where("product_id=?", p.ProductId).Update("is_deleted", p.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
