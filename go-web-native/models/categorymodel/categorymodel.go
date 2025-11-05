package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`Select * From categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdateAt); err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories

}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`Insert into categories (name, created_at, updated_at) value (?,?,?)`, category.Name, category.CreatedAt, category.UpdateAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInsertId > 0
}

func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`select id, name from categories where id = ?`, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`Update categories set name = ?, updated_at = ? Where id = ?`, category.Name, category.UpdateAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0

}

func Delete(id int) error {
	_, err := config.DB.Exec(`delete from categories where id = ?`, id)
	return err
}
