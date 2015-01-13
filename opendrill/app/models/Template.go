package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Template struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `json:"name"`
	DesignerID bson.ObjectId `json:"designerId" bson:"designerId"`
	Image      string        `json:"image"`
	Content    string        `json:"content"`
	Type       string        `json:"type"`
	Selected   string        `json:"selected"`
	CreatedAt  time.Time     `json:"createdAt"`
	ModifiedAt time.Time     `json:"updatedAt"`
}

func GetTemplatesFromDesigner(designerID string) (templates2 []Template) {
	templates.
		Find(bson.M{"designerId": bson.ObjectIdHex(designerID)}).
		All(&templates2)
	return
}

func GetTemplateFromDesigner(designerID string, templateID string) (err error, template2 Template) {
	err = templates.
		Find(bson.M{"designerId": bson.ObjectIdHex(designerID), "_id": bson.ObjectIdHex(templateID)}).
		One(&template2)
	return
}

func AddTemplateToDesigner(template Template, designerID string) (err error, template2 Template) {
	var designer Designer
	err, designer = GetDesigner(designerID)
	if err != nil {
		return err, template
	}
	template2 = template
	template2.Id = bson.NewObjectId()
	template2.DesignerID = designer.Id
	if err := templates.Insert(template2); err != nil {
		return err, template
	}

	designers.Update(bson.M{"_id": designer.Id},
		bson.M{
			"templates": GetTemplatesFromDesigner(designerID),
		})
	return nil, template2
}

func RemoveTemplateFromDesigner(designerID string, templateID string) (err error, deleted bool) {
	// deleted = false
	// bid := bson.ObjectIdHex(Id)
	// err = designers.
	// 	Remove(bson.M{"_id": bid})
	// if err != nil {
	// 	return err, deleted
	// }
	deleted = true
	return nil, deleted
}

func UpdateTemplateFromDesigner(template Template, designerID string, templateID string) (err error, template2 Template) {
	var designer Designer
	err, designer = GetDesigner(designerID)
	if err != nil {
		return err, template
	}

	template2 = template
	bid := bson.ObjectIdHex(templateID)
	err = templates.Update(bson.M{"_id": bid},
		bson.M{"name": template2.Name,
			"image":      template2.Image,
			"content":    template2.Content,
			"type":       template2.Type,
			"_id":        bid,
			"designerId": bson.ObjectIdHex(designerID),
			"CreatedAt":  template2.CreatedAt,
			"ModifiedAt": time.Now(),
		})
	if err != nil {
		return err, template
	}

	designers.Update(bson.M{"_id": designer.Id},
		bson.M{
			"templates": GetTemplatesFromDesigner(designerID),
		})
	return nil, template2
}
