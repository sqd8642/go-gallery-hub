# go-gallery-hub

## Overview
Go App Dev 2024 Spring semseter course project. Image gallery application built with Go programming language.
It offers a set of endpoints to manage images, galleries, and tags. Below are the supported operations

```
Create Gallery: POST /galleries
Get Gallery by ID: GET /galleries/:id
Update Gallery: PUT /galleries/:id
Delete Gallery: DELETE /galleries/:id
Create Image: POST /galleries/:id/images
Get Images in Gallery: GET /galleries/:id/images
Create Tag: POST /tags
Get Tag by ID: GET /tags/:id
Update Tag: PUT /tags/:id
Delete Tag: DELETE /tags/:id
```

## Database structure

```
Table galleries {
  id bigserial [primary key]
  created_at timestamp
  updated_at timestamp
  title text
  description text
}

Table images {
  id bigserial [primary key]
  created_at timestamp
  updated_at timestamp
  gallery bigserial
  url text
  caption text
}

Table galleries_and_images {
  id bigserial [primary key]
  created_at timestamp
  updated_at timestamp
  gallery bigserial
  image bigserial
}

Ref: galleries_and_images.restaurant < galleries.id
Ref: galleries_and_images.menu < image.id
```

