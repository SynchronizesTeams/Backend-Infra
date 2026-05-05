# 🚀 API Documentation

## 1. Project Overview

This is the official API documentation for the **Backend-Infra** project, serving as the core infrastructure backend.

**Tech Stack:**
- **Language:** Go 1.21+
- **Framework:** Fiber (v2)
- **ORM:** GORM
- **Database:** MySQL
- **Authentication:** JWT (JSON Web Tokens)

## 4. Authentication & Authorization

The API uses **JWT (JSON Web Token)** for authentication.
1. **Login Flow:** Send a `POST` request to `/api/v1/auth/login` with `email` and `password`.
2. **Token Format:** The server responds with a token. Include it in the `Authorization` header as `Bearer <token>` for protected routes.

Example Header:
```http
Authorization: Bearer eyJhbGciOiJIUzI1...
```

## 5. Error Handling

Errors are standardized in JSON format with appropriate HTTP status codes (400, 401, 403, 404, 500).

```json
{
  "error": "Detailed error message"
}
```

## 7. Setup & Usage

1. Clone the repository.
2. Copy `.env.example` to `.env` and fill the variables (DB credentials, `JWT_SECRET`, `APP_PORT`).
3. Run `go run main.go`.
4. GORM will auto-migrate the database schemas on startup.

## 2. Database Schema

### `Certification`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Name | `string` | `-` | `` |
| Issuer | `string` | `-` | `` |
| Description | `string` | `-` | `` |
| Image | `string` | `-` | `` |
| CertificationNumber | `string` | `-` | `` |
| IssueDate | `string` | `-` | `` |
| ExpiryDate | `string` | `-` | `` |


### `UserLinks`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Title | `string` | `title` | `` |
| Url | `string` | `url` | `` |
| Icon | `string` | `icon` | `` |
| UserID | `uint` | `user_id` | `` |
| User | `User` | `user` | `constraint:OnUpdate:CASCADE,OnDelete:CASCADE;` |


### `ForumSection`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Name | `string` | `-` | `` |
| Slug | `string` | `-` | `` |
| Description | `string` | `-` | `type:text` |
| Icon | `string` | `-` | `` |
| IsActive | `bool` | `-` | `default:true` |


### `Testimonial`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Name | `string` | `-` | `` |
| Position | `string` | `-` | `` |
| Photo | `string` | `-` | `` |
| Testimonial | `string` | `-` | `` |


### `GuestBook`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Title | `string` | `title` | `` |
| InstanceName | `string` | `instance_name` | `` |
| ContactPerson | `string` | `contact_person` | `` |
| Email | `string` | `email` | `` |
| Phone | `string` | `phone` | `` |
| Description | `string` | `description` | `` |
| Status | `string` | `status` | `type:enum('pending','approved','rejected','hidden');default:'pending'` |
| RejectionReason | `string` | `rejection_reason` | `type:text` |
| ApprovedAt | `time.Time` | `approved_at` | `` |
| RequestDate | `time.Time` | `request_date` | `type:date` |
| ShowInCalendar | `bool` | `show_in_calendar` | `` |


### `Event`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Title | `string` | `title` | `` |
| Description | `string` | `description` | `type:text` |
| Category | `string` | `category` | `type:enum('kunjungan_industri', 'tamu', 'acara_sekolah', 'ujian', 'libur', 'lainnya');default:'tamu'` |
| Visibility | `string` | `visibility` | `type:enum('public', 'private');default:'private'` |
| StartDate | `string` | `start_date` | `` |
| EndDate | `string` | `end_date` | `` |
| Location | `string` | `location` | `` |
| Organizer | `string` | `organizer` | `` |
| Image | `string` | `image` | `` |


### `Achievement`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Title | `string` | `title` | `` |
| Description | `string` | `description` | `` |
| Image | `string` | `image` | `` |
| AchievementDate | `string` | `achievement_date` | `` |
| NewsID | `uint` | `news_id` | `` |
| News | `News` | `news` | `` |


### `Industry`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Name | `string` | `-` | `` |
| Logo | `string` | `-` | `` |
| Website | `string` | `-` | `` |
| Description | `string` | `-` | `type:text` |


### `Image`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| FilePath | `string` | `file_path` | `` |
| Category | `string` | `category` | `type:enum('berita', 'mading', 'galeri', 'teacher', 'eskul', 'achievement', 'testimonial', 'saprol', 'certification', 'portal', 'mitra', 'profile');default:'galeri'` |
| Title | `string` | `title` | `` |
| EventID | `uint` | `event` | `` |


### `Teacher`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| NIG | `string` | `nig` | `` |
| FullName | `string` | `full_name` | `` |
| Position | `string` | `position` | `` |
| Subject | `string` | `subject` | `` |
| Photo | `string` | `photo` | `` |
| Description | `string` | `description` | `` |
| UserID | `uint` | `user_id` | `` |
| User | `User` | `-` | `` |
| Eskuls | `[]Eskul` | `eskuls` | `foreignKey:PembinaID` |


### `NewsTag`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Name | `string` | `name` | `unique;not null` |
| News | `[]News` | `news` | `many2many:news_news_tags;` |


### `ForumReply`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Content | `string` | `content` | `type:text; not null` |
| UserID | `uint` | `user_id` | `` |
| PostID | `uint` | `post_id` | `` |
| ParentID | `*uint` | `parent_id` | `` |
| User | `User` | `user` | `` |
| Post | `ForumPost` | `post` | `` |
| Parent | `*ForumReply` | `parent` | `foreignKey:ParentID` |
| Children | `[]ForumReply` | `children` | `foreignKey:ParentID;constraint:OnDelete:CASCADE` |


### `News`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Title | `string` | `title` | `` |
| Slug | `string` | `slug` | `` |
| Content | `string` | `content` | `type:text` |
| Excerpt | `string` | `excerpt` | `` |
| Thumbnail | `string` | `thumbnail` | `` |
| Status | `string` | `status` | `type:enum('draft', 'published', 'archived');default:'draft'` |
| ViewCount | `int` | `view_count` | `` |
| Tags | `[]NewsTag` | `tags` | `many2many:news_news_tags;` |
| AuthorID | `uint` | `author_id` | `` |
| Author | `User` | `author` | `constraint:OnUpdate:CASCADE,OnDelete:CASCADE;` |


### `ForumPost`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Title | `string` | `title` | `not null` |
| Content | `string` | `content` | `type:text` |
| Image | `string` | `image` | `` |
| File | `string` | `file` | `` |
| Status | `string` | `-` | `type:enum('pending','approved','rejected','hidden');default:'pending'` |
| Upvote | `int` | `upvote` | `` |
| Downvote | `int` | `downvote` | `` |
| ReplyCount | `int` | `reply_count` | `` |
| IsHidden | `bool` | `is_hidden` | `` |
| UserID | `uint` | `user_id` | `` |
| SectionID | `uint` | `section_id` | `` |
| User | `User` | `user` | `` |
| Section | `ForumSection` | `section` | `` |


### `User`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| NoInduk | `string` | `no_induk` | `` |
| Name | `string` | `name` | `` |
| Email | `string` | `email` | `unique` |
| Password | `string` | `-` | `` |
| Role | `string` | `role` | `type:enum('admin','guru','siswa','orang_tua');default:'siswa'` |
| PhotoUrl | `string` | `photo_url` | `` |
| Phone | `string` | `phone` | `` |
| Alamat | `string` | `alamat` | `type:text` |
| Jabatan | `string` | `jabatan` | `` |
| TahunAjaranMulai | `string` | `tahun_ajaran_mulai` | `type:year;default:null` |
| UserLinks | `[]UserLinks` | `user_links` | `foreignKey:UserID` |
| News | `[]News` | `news` | `foreignKey:AuthorID` |
| Post | `[]ForumPost` | `-` | `foreignKey:UserID` |
| Replies | `[]ForumReply` | `-` | `foreignKey:UserID` |
| Teacher | `[]Teacher` | `-` | `foreignKey:UserID` |


### `Portal`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Name | `string` | `-` | `` |
| Description | `string` | `-` | `` |
| URL | `string` | `-` | `` |
| Logo | `string` | `-` | `` |
| Category | `string` | `-` | `` |
| IsSSOEnabled | `bool` | `-` | `default:false` |
| IsActive | `bool` | `-` | `default:true` |


### `Eskul`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Name | `string` | `name` | `` |
| Description | `string` | `description` | `type:text` |
| Image | `string` | `image` | `` |
| PembinaID | `uint` | `pembina_id` | `` |
| Pembina | `Teacher` | `pembina` | `foreignKey:PembinaID` |


### `Mading`
| Field Name | Go Type | JSON Key | Constraints / GORM Tag |
|---|---|---|---|
| Title | `string` | `title` | `` |
| Type | `string` | `type` | `type:enum('infographic', 'announcement');default:'infographic'` |
| Content | `string` | `content` | `type:text` |
| Image | `string` | `image` | `` |
| Status | `string` | `status` | `type:enum('draft', 'published', 'archieve');default:'draft'` |
| IsActive | `bool` | `is_active` | `` |


## 3. API Endpoints

### /SHOW Routes

#### `GET /show/:id`
**Description**: Handled by `ShowAchievement`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/:id`
**Description**: Handled by `ShowCertification`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/:id`
**Description**: Handled by `ShowEskul`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/:id`
**Description**: Handled by `GetEvent`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/date/:date`
**Description**: Handled by `GetEventByDate`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/:id`
**Description**: Handled by `GetGuestBook`
**Authentication**: None

**Response Format**: JSON

#### `GET /show`
**Description**: Handled by `ShowImagesByCategory`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/:id`
**Description**: Handled by `ShowIndustry`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/:id`
**Description**: Handled by `ShowMading`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/:id`
**Description**: Handled by `ShowPortal`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/:id`
**Description**: Handled by `ShowTeacher`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/:id`
**Description**: Handled by `ShowTestimonial`
**Authentication**: None

**Response Format**: JSON

#### `GET /show/:id`
**Description**: Handled by `ShowUserLinks`
**Authentication**: None

**Response Format**: JSON

### /SHOWALL Routes

#### `GET /showAll`
**Description**: Handled by `ShowAllAchievement`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `ShowAllCertification`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `ShowAllEskul`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `GetAllEvent`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `GetAllGuestBook`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `ShowAllImages`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `ShowAllIndustry`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `ShowAllMadings`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `ShowAllPortal`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `ShowAllTeachers`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `ShowAllTestimonial`
**Authentication**: None

**Response Format**: JSON

#### `GET /showAll`
**Description**: Handled by `ShowAllSection`
**Authentication**: None

**Response Format**: JSON

### /CREATE Routes

#### `POST /create`
**Description**: Handled by `CreateGuest`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "instance_name": "string",
  "contact_person": "string",
  "email": "string",
  "phone": "string",
  "description": "string",
  "request_date": "string",
}
```

**Response Format**: JSON

### /GETALL Routes

#### `GET /getAll`
**Description**: Handled by `GetAllNews`
**Authentication**: None

**Response Format**: JSON

### /GET Routes

#### `GET /get/:id`
**Description**: Handled by `GetNews`
**Authentication**: None

**Response Format**: JSON

### /:ID Routes

#### `GET profile/:id`
**Description**: Handled by `ProfilePublic`
**Authentication**: None

**Response Format**: JSON

### /CERTIFICATION Routes

#### `POST /certification/create`
**Description**: Handled by `CreateCertification`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "issuer": "string",
  "description": "string",
  "image": "string",
  "certification_number": "string",
  "issue_date": "string",
  "expiry_date": "string",
}
```

**Response Format**: JSON

#### `POST /certification/edit/:id`
**Description**: Handled by `EditCertification`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "issuer": "string",
  "description": "string",
  "image": "string",
  "certification_number": "string",
  "issue_date": "string",
  "expiry_date": "string",
}
```

**Response Format**: JSON

#### `GET /certification/show/:id`
**Description**: Handled by `ShowCertification`
**Authentication**: None

**Response Format**: JSON

#### `GET /certification/showAll`
**Description**: Handled by `ShowAllCertification`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /certification/delete/:id`
**Description**: Handled by `DeleteCertification`
**Authentication**: None

**Response Format**: JSON

### /AUTH Routes

#### `POST /auth/register`
**Description**: Handled by `Register`
**Authentication**: Required (JWT)

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "email": "string",
  "password": "string",
}
```

**Response Format**: JSON

#### `POST /auth/login`
**Description**: Handled by `Login`
**Authentication**: Required (JWT)

**Request Body Schema** (JSON):
```json
{
  "email": "string",
  "password": "string",
}
```

**Response Format**: JSON

#### `GET /auth/profile`
**Description**: Handled by `Profile`
**Authentication**: Required (JWT)

**Response Format**: JSON

#### `POST /auth/edit`
**Description**: Handled by `EditProfile`
**Authentication**: Required (JWT)

**Request Body Schema** (JSON):
```json
{
  "no_induk": "*string",
  "name": "*string",
  "email": "*string",
  "role": "*string",
  "photo_url": "*string",
  "phone": "*string",
  "alamat": "*string",
  "jabatan": "*string",
  "tahun_ajaran_mulai": "*string",
}
```

**Response Format**: JSON

### /FORUM-SECTION Routes

#### `POST /forum-section/create`
**Description**: Handled by `CreateSection`
**Authentication**: None

**Response Format**: JSON

#### `POST /forum-section/edit/:id`
**Description**: Handled by `EditSection`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /forum-section/delete/:id`
**Description**: Handled by `DeleteSection`
**Authentication**: None

**Response Format**: JSON

#### `GET /forum-section/showAll`
**Description**: Handled by `ShowAllSection`
**Authentication**: None

**Response Format**: JSON

### /TESTIMONIAL Routes

#### `POST /testimonial/create`
**Description**: Handled by `CreateTestimonial`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "position": "string",
  "testimonial": "string",
}
```

**Response Format**: JSON

#### `POST /testimonial/edit/:id`
**Description**: Handled by `EditTestimonial`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "position": "string",
  "testimonial": "string",
}
```

**Response Format**: JSON

#### `GET /testimonial/show/:id`
**Description**: Handled by `ShowTestimonial`
**Authentication**: None

**Response Format**: JSON

#### `GET /testimonial/showAll`
**Description**: Handled by `ShowAllTestimonial`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /testimonial/delete/:id`
**Description**: Handled by `DeleteTestimonial`
**Authentication**: None

**Response Format**: JSON

### /EVENT Routes

#### `POST /event/create`
**Description**: Handled by `CreateEvent`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "description": "string",
  "category": "string",
  "start_date": "string",
  "end_date": "string",
  "location": "string",
  "organizer": "string",
}
```

**Response Format**: JSON

#### `POST /event/edit/:id`
**Description**: Handled by `EditEvent`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "description": "string",
  "category": "string",
  "start_date": "string",
  "end_date": "string",
  "location": "string",
  "organizer": "string",
}
```

**Response Format**: JSON

#### `DELETE /event/delete/:id`
**Description**: Handled by `DeleteEvent`
**Authentication**: None

**Response Format**: JSON

#### `GET /event/show/:id`
**Description**: Handled by `GetEvent`
**Authentication**: None

**Response Format**: JSON

#### `GET /event/showAll`
**Description**: Handled by `GetAllEvent`
**Authentication**: None

**Response Format**: JSON

#### `PUT /event/change-visibility/:id`
**Description**: Handled by `ChangeVisibility`
**Authentication**: None

**Response Format**: JSON

### /ACHIEVEMENT Routes

#### `POST /achievement/create`
**Description**: Handled by `CreateAchievement`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "description": "string",
  "image": "string",
  "achievement_date": "string",
  "news_id": "uint",
}
```

**Response Format**: JSON

#### `POST /achievement/edit/:id`
**Description**: Handled by `EditAchievement`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "description": "string",
  "image": "string",
  "achievement_date": "string",
  "news_id": "uint",
}
```

**Response Format**: JSON

#### `GET /achievement/show/:id`
**Description**: Handled by `ShowAchievement`
**Authentication**: None

**Response Format**: JSON

#### `GET /achievement/showAll`
**Description**: Handled by `ShowAllAchievement`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /achievement/delete/:id`
**Description**: Handled by `DeleteAchievement`
**Authentication**: None

**Response Format**: JSON

### /INDUSTRY Routes

#### `POST /industry/create`
**Description**: Handled by `CreateIndustry`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "logo": "string",
  "website": "string",
  "description": "string",
}
```

**Response Format**: JSON

#### `POST /industry/edit/:id`
**Description**: Handled by `EditIndustry`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "logo": "string",
  "website": "string",
  "description": "string",
}
```

**Response Format**: JSON

#### `GET /industry/show/:id`
**Description**: Handled by `ShowIndustry`
**Authentication**: None

**Response Format**: JSON

#### `GET /industry/showAll`
**Description**: Handled by `ShowAllIndustry`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /industry/delete/:id`
**Description**: Handled by `DeleteIndustry`
**Authentication**: None

**Response Format**: JSON

### /NEWS Routes

#### `POST /news/create`
**Description**: Handled by `CreateNews`
**Authentication**: None

**Response Format**: JSON

#### `POST /news/update/:id`
**Description**: Handled by `UpdateNews`
**Authentication**: None

**Response Format**: JSON

#### `GET /news/getAll`
**Description**: Handled by `GetAllNews`
**Authentication**: None

**Response Format**: JSON

#### `GET /news/get/:id`
**Description**: Handled by `GetNews`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /news/delete/:id`
**Description**: Handled by `DeleteNews`
**Authentication**: None

**Response Format**: JSON

#### `POST /news/addView/:id`
**Description**: Handled by `AddView`
**Authentication**: None

**Response Format**: JSON

#### `POST /news/changeStatus/:id`
**Description**: Handled by `ChangeStatus`
**Authentication**: None

**Response Format**: JSON

### /TEACHER Routes

#### `POST /teacher/create`
**Description**: Handled by `CreateTeacher`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "nig": "string",
  "full_name": "string",
  "position": "string",
  "subject": "string",
  "description": "string",
}
```

**Response Format**: JSON

#### `POST /teacher/edit/:id`
**Description**: Handled by `EditTeacher`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "nig": "string",
  "full_name": "string",
  "position": "string",
  "subject": "string",
  "description": "string",
}
```

**Response Format**: JSON

#### `DELETE /teacher/delete/:id`
**Description**: Handled by `DeleteTeacher`
**Authentication**: None

**Response Format**: JSON

#### `GET /teacher/showAll`
**Description**: Handled by `ShowAllTeachers`
**Authentication**: None

**Response Format**: JSON

#### `GET /teacher/show/:id`
**Description**: Handled by `ShowTeacher`
**Authentication**: None

**Response Format**: JSON

### /USER-LINKS Routes

#### `POST /user-links/add`
**Description**: Handled by `AddLink`
**Authentication**: None

**Response Format**: JSON

#### `POST /user-links/edit/:id`
**Description**: Handled by `EditLinks`
**Authentication**: None

**Response Format**: JSON

#### `GET /user-links/show/self`
**Description**: Handled by `ShowUserLinksSelf`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /user-links/delete/:id`
**Description**: Handled by `DeleteLinks`
**Authentication**: None

**Response Format**: JSON

### /GUEST-BOOK Routes

#### `POST /guest-book/create`
**Description**: Handled by `CreateGuest`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "instance_name": "string",
  "contact_person": "string",
  "email": "string",
  "phone": "string",
  "description": "string",
  "request_date": "string",
}
```

**Response Format**: JSON

#### `POST /guest-book/edit/:id`
**Description**: Handled by `EditGuest`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "instance_name": "string",
  "contact_person": "string",
  "email": "string",
  "phone": "string",
  "description": "string",
  "request_date": "string",
}
```

**Response Format**: JSON

#### `GET /guest-book/showAll`
**Description**: Handled by `GetAllGuestBook`
**Authentication**: None

**Response Format**: JSON

#### `GET /guest-book/show/:id`
**Description**: Handled by `GetGuestBook`
**Authentication**: None

**Response Format**: JSON

#### `PUT /guest-book/approve/:id`
**Description**: Handled by `ApprovedGuest`
**Authentication**: None

**Response Format**: JSON

#### `PUT /guest-book/reject/:id`
**Description**: Handled by `RejectGuestBook`
**Authentication**: None

**Response Format**: JSON

### /FORUM-POST Routes

#### `POST /forum-post/create`
**Description**: Handled by `CreatePost`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "content": "string",
  "section_id": "uint",
}
```

**Response Format**: JSON

#### `POST /forum-post/edit/:id`
**Description**: Handled by `EditPost`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "content": "string",
  "section_id": "uint",
}
```

**Response Format**: JSON

#### `DELETE /forum-post/delete/:id`
**Description**: Handled by `DeletePost`
**Authentication**: None

**Response Format**: JSON

#### `GET /forum-post/showAll`
**Description**: Handled by `GetAllPost`
**Authentication**: None

**Response Format**: JSON

#### `GET /forum-post/show/:id`
**Description**: Handled by `ShowPostById`
**Authentication**: None

**Response Format**: JSON

#### `PUT /forum-post/changeStatus`
**Description**: Handled by `ChangePostStatus`
**Authentication**: None

**Response Format**: JSON

#### `PUT /forum-post/upvote/:id`
**Description**: Handled by `UpvotePost`
**Authentication**: None

**Response Format**: JSON

#### `PUT /forum-post/downvote/:id`
**Description**: Handled by `DownvotePost`
**Authentication**: None

**Response Format**: JSON

#### `POST /forum-post/reply/:id`
**Description**: Handled by `CreateReply`
**Authentication**: None

**Response Format**: JSON

#### `GET /forum-post/post-replies/:id`
**Description**: Handled by `GetPostWithReplies`
**Authentication**: None

**Response Format**: JSON

#### `GET /forum-post/replies/:id`
**Description**: Handled by `GetNestedReplies`
**Authentication**: None

**Response Format**: JSON

### /PORTAL Routes

#### `POST /portal/create`
**Description**: Handled by `CreatePortal`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "description": "string",
  "url": "string",
  "logo": "string",
  "category": "string",
}
```

**Response Format**: JSON

#### `POST /portal/edit/:id`
**Description**: Handled by `EditPortal`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "description": "string",
  "url": "string",
  "logo": "string",
  "category": "string",
}
```

**Response Format**: JSON

#### `GET /portal/show/:id`
**Description**: Handled by `ShowPortal`
**Authentication**: None

**Response Format**: JSON

#### `GET /portal/showAll`
**Description**: Handled by `ShowAllPortal`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /portal/delete/:id`
**Description**: Handled by `DeletePortal`
**Authentication**: None

**Response Format**: JSON

### /ESKUL Routes

#### `POST /eskul/create`
**Description**: Handled by `CreateEskul`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "description": "string",
  "pembina_id": "uint",
}
```

**Response Format**: JSON

#### `POST /eskul/edit/:id`
**Description**: Handled by `EditEskul`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "name": "string",
  "description": "string",
  "pembina_id": "uint",
}
```

**Response Format**: JSON

#### `GET /eskul/show/:id`
**Description**: Handled by `ShowEskul`
**Authentication**: None

**Response Format**: JSON

#### `GET /eskul/showAll`
**Description**: Handled by `ShowAllEskul`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /eskul/delete/:id`
**Description**: Handled by `DeleteEskul`
**Authentication**: None

**Response Format**: JSON

### /IMAGE Routes

#### `POST /image/add`
**Description**: Handled by `AddImage`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /image/delete/:id`
**Description**: Handled by `DeleteImage`
**Authentication**: None

**Response Format**: JSON

#### `GET /image/showAll`
**Description**: Handled by `ShowAllImages`
**Authentication**: None

**Response Format**: JSON

#### `GET /image/show`
**Description**: Handled by `ShowImagesByCategory`
**Authentication**: None

**Response Format**: JSON

### /MADING Routes

#### `POST /mading/create`
**Description**: Handled by `CreateMading`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "type": "string",
  "content": "string",
  "image": "uint",
}
```

**Response Format**: JSON

#### `POST /mading/edit/:id`
**Description**: Handled by `EditMading`
**Authentication**: None

**Request Body Schema** (JSON):
```json
{
  "title": "string",
  "type": "string",
  "content": "string",
  "image": "uint",
}
```

**Response Format**: JSON

#### `GET /mading/show/:id`
**Description**: Handled by `ShowMading`
**Authentication**: None

**Response Format**: JSON

#### `GET /mading/showAll`
**Description**: Handled by `ShowAllMadings`
**Authentication**: None

**Response Format**: JSON

#### `DELETE /mading/delete/:id`
**Description**: Handled by `DeleteMading`
**Authentication**: None

**Response Format**: JSON

### /MADINGCHANGE-STATUS Routes

#### `PUT /madingchange-status/:id`
**Description**: Handled by `ChangeStatusMading`
**Authentication**: None

**Response Format**: JSON
