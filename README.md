# go-gallery-hub

## Overview
Go App Dev 2024 Spring semseter course project. Image gallery application built with Go programming language.
It offers a set of endpoints to manage images, galleries, and tags. Below are the supported operations

```
Create Image: POST /images
Get Image by ID: GET /images/:id
Update Image: PUT /images/:id
Delete Image: DELETE images/:id
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

