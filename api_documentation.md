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

### 📁 Public Routes
Kumpulan endpoint untuk modul **Public**.

#### `GET /public/achievement/show/:id`
**Group Route:** `/public/achievement`
**Handler:** `ShowAchievement`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/achievement/showAll`
**Group Route:** `/public/achievement`
**Handler:** `ShowAllAchievement`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/certification/show/:id`
**Group Route:** `/public/certification`
**Handler:** `ShowCertification`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/certification/showAll`
**Group Route:** `/public/certification`
**Handler:** `ShowAllCertification`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/eskul/show/:id`
**Group Route:** `/public/eskul`
**Handler:** `ShowEskul`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/eskul/showAll`
**Group Route:** `/public/eskul`
**Handler:** `ShowAllEskul`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/event/show/:id`
**Group Route:** `/public/event`
**Handler:** `GetEvent`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/event/showAll`
**Group Route:** `/public/event`
**Handler:** `GetAllEvent`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/event/show/date/:date`
**Group Route:** `/public/event`
**Handler:** `GetEventByDate`
**Authentication:** None

**Response Format**: JSON

#### `POST /public/guest/create`
**Group Route:** `/public/guest`
**Handler:** `CreateGuest`
**Authentication:** Required (JWT)

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

#### `GET /public/guest/showAll`
**Group Route:** `/public/guest`
**Handler:** `GetAllGuestBook`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/guest/show/:id`
**Group Route:** `/public/guest`
**Handler:** `GetGuestBook`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/image/showAll`
**Group Route:** `/public/image`
**Handler:** `ShowAllImages`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/image/show`
**Group Route:** `/public/image`
**Handler:** `ShowImagesByCategory`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/industry/show/:id`
**Group Route:** `/public/industry`
**Handler:** `ShowIndustry`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/industry/showAll`
**Group Route:** `/public/industry`
**Handler:** `ShowAllIndustry`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/mading/show/:id`
**Group Route:** `/public/mading`
**Handler:** `ShowMading`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/mading/showAll`
**Group Route:** `/public/mading`
**Handler:** `ShowAllMadings`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/news/getAll`
**Group Route:** `/public/news`
**Handler:** `GetAllNews`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/news/get/:id`
**Group Route:** `/public/news`
**Handler:** `GetNews`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/portal/show/:id`
**Group Route:** `/public/portal`
**Handler:** `ShowPortal`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/portal/showAll`
**Group Route:** `/public/portal`
**Handler:** `ShowAllPortal`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/teacher/showAll`
**Group Route:** `/public/teacher`
**Handler:** `ShowAllTeachers`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/teacher/show/:id`
**Group Route:** `/public/teacher`
**Handler:** `ShowTeacher`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/testimonial/showAll`
**Group Route:** `/public/testimonial`
**Handler:** `ShowAllTestimonial`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/testimonial/show/:id`
**Group Route:** `/public/testimonial`
**Handler:** `ShowTestimonial`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/user-link/show/:id`
**Group Route:** `/public/user-link`
**Handler:** `ShowUserLinks`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/forum-section/showAll`
**Group Route:** `/public/forum-section`
**Handler:** `ShowAllSection`
**Authentication:** None

**Response Format**: JSON

#### `GET /public/userprofile/:id`
**Group Route:** `/public/user`
**Handler:** `ProfilePublic`
**Authentication:** None

**Response Format**: JSON

### 📁 Certification Routes
Kumpulan endpoint untuk modul **Certification**.

#### `POST /certification/create`
**Group Route:** `/certification`
**Handler:** `CreateCertification`
**Authentication:** Required (JWT)

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
**Group Route:** `/certification`
**Handler:** `EditCertification`
**Authentication:** Required (JWT)

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
**Group Route:** `/certification`
**Handler:** `ShowCertification`
**Authentication:** None

**Response Format**: JSON

#### `GET /certification/showAll`
**Group Route:** `/certification`
**Handler:** `ShowAllCertification`
**Authentication:** None

**Response Format**: JSON

#### `DELETE /certification/delete/:id`
**Group Route:** `/certification`
**Handler:** `DeleteCertification`
**Authentication:** Required (JWT)

**Response Format**: JSON

### 📁 Auth Routes
Kumpulan endpoint untuk modul **Auth**.

#### `POST /auth/register`
**Group Route:** `/auth`
**Handler:** `Register`
**Authentication:** Required (JWT)

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
**Group Route:** `/auth`
**Handler:** `Login`
**Authentication:** Required (JWT)

**Request Body Schema** (JSON):
```json
{
  "email": "string",
  "password": "string",
}
```

**Response Format**: JSON

#### `GET /auth/profile`
**Group Route:** `/auth`
**Handler:** `Profile`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `POST /auth/edit`
**Group Route:** `/auth`
**Handler:** `EditProfile`
**Authentication:** Required (JWT)

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

### 📁 Forum_section Routes
Kumpulan endpoint untuk modul **Forum_section**.

#### `POST /forum-section/create`
**Group Route:** `/forum-section`
**Handler:** `CreateSection`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `POST /forum-section/edit/:id`
**Group Route:** `/forum-section`
**Handler:** `EditSection`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `DELETE /forum-section/delete/:id`
**Group Route:** `/forum-section`
**Handler:** `DeleteSection`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `GET /forum-section/showAll`
**Group Route:** `/forum-section`
**Handler:** `ShowAllSection`
**Authentication:** None

**Response Format**: JSON

### 📁 Testimonial Routes
Kumpulan endpoint untuk modul **Testimonial**.

#### `POST /testimonial/create`
**Group Route:** `/testimonial`
**Handler:** `CreateTestimonial`
**Authentication:** Required (JWT)

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
**Group Route:** `/testimonial`
**Handler:** `EditTestimonial`
**Authentication:** Required (JWT)

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
**Group Route:** `/testimonial`
**Handler:** `ShowTestimonial`
**Authentication:** None

**Response Format**: JSON

#### `GET /testimonial/showAll`
**Group Route:** `/testimonial`
**Handler:** `ShowAllTestimonial`
**Authentication:** None

**Response Format**: JSON

#### `DELETE /testimonial/delete/:id`
**Group Route:** `/testimonial`
**Handler:** `DeleteTestimonial`
**Authentication:** Required (JWT)

**Response Format**: JSON

### 📁 Event Routes
Kumpulan endpoint untuk modul **Event**.

#### `POST /event/create`
**Group Route:** `/event`
**Handler:** `CreateEvent`
**Authentication:** Required (JWT)

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
**Group Route:** `/event`
**Handler:** `EditEvent`
**Authentication:** Required (JWT)

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
**Group Route:** `/event`
**Handler:** `DeleteEvent`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `GET /event/show/:id`
**Group Route:** `/event`
**Handler:** `GetEvent`
**Authentication:** None

**Response Format**: JSON

#### `GET /event/showAll`
**Group Route:** `/event`
**Handler:** `GetAllEvent`
**Authentication:** None

**Response Format**: JSON

#### `PUT /event/change-visibility/:id`
**Group Route:** `/event`
**Handler:** `ChangeVisibility`
**Authentication:** None

**Response Format**: JSON

### 📁 Achievement Routes
Kumpulan endpoint untuk modul **Achievement**.

#### `POST /achievement/create`
**Group Route:** `/achievement`
**Handler:** `CreateAchievement`
**Authentication:** Required (JWT)

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
**Group Route:** `/achievement`
**Handler:** `EditAchievement`
**Authentication:** Required (JWT)

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
**Group Route:** `/achievement`
**Handler:** `ShowAchievement`
**Authentication:** None

**Response Format**: JSON

#### `GET /achievement/showAll`
**Group Route:** `/achievement`
**Handler:** `ShowAllAchievement`
**Authentication:** None

**Response Format**: JSON

#### `DELETE /achievement/delete/:id`
**Group Route:** `/achievement`
**Handler:** `DeleteAchievement`
**Authentication:** Required (JWT)

**Response Format**: JSON

### 📁 Industry Routes
Kumpulan endpoint untuk modul **Industry**.

#### `POST /industry/create`
**Group Route:** `/industry`
**Handler:** `CreateIndustry`
**Authentication:** Required (JWT)

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
**Group Route:** `/industry`
**Handler:** `EditIndustry`
**Authentication:** Required (JWT)

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
**Group Route:** `/industry`
**Handler:** `ShowIndustry`
**Authentication:** None

**Response Format**: JSON

#### `GET /industry/showAll`
**Group Route:** `/industry`
**Handler:** `ShowAllIndustry`
**Authentication:** None

**Response Format**: JSON

#### `DELETE /industry/delete/:id`
**Group Route:** `/industry`
**Handler:** `DeleteIndustry`
**Authentication:** Required (JWT)

**Response Format**: JSON

### 📁 News Routes
Kumpulan endpoint untuk modul **News**.

#### `POST /news/create`
**Group Route:** `/news`
**Handler:** `CreateNews`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `POST /news/update/:id`
**Group Route:** `/news`
**Handler:** `UpdateNews`
**Authentication:** None

**Response Format**: JSON

#### `GET /news/getAll`
**Group Route:** `/news`
**Handler:** `GetAllNews`
**Authentication:** None

**Response Format**: JSON

#### `GET /news/get/:id`
**Group Route:** `/news`
**Handler:** `GetNews`
**Authentication:** None

**Response Format**: JSON

#### `DELETE /news/delete/:id`
**Group Route:** `/news`
**Handler:** `DeleteNews`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `POST /news/addView/:id`
**Group Route:** `/news`
**Handler:** `AddView`
**Authentication:** None

**Response Format**: JSON

#### `POST /news/changeStatus/:id`
**Group Route:** `/news`
**Handler:** `ChangeStatus`
**Authentication:** None

**Response Format**: JSON

### 📁 Teacher Routes
Kumpulan endpoint untuk modul **Teacher**.

#### `POST /teacher/create`
**Group Route:** `/teacher`
**Handler:** `CreateTeacher`
**Authentication:** Required (JWT)

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
**Group Route:** `/teacher`
**Handler:** `EditTeacher`
**Authentication:** Required (JWT)

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
**Group Route:** `/teacher`
**Handler:** `DeleteTeacher`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `GET /teacher/showAll`
**Group Route:** `/teacher`
**Handler:** `ShowAllTeachers`
**Authentication:** None

**Response Format**: JSON

#### `GET /teacher/show/:id`
**Group Route:** `/teacher`
**Handler:** `ShowTeacher`
**Authentication:** None

**Response Format**: JSON

### 📁 User_links Routes
Kumpulan endpoint untuk modul **User_links**.

#### `POST /user-links/add`
**Group Route:** `/user-links`
**Handler:** `AddLink`
**Authentication:** None

**Response Format**: JSON

#### `POST /user-links/edit/:id`
**Group Route:** `/user-links`
**Handler:** `EditLinks`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `GET /user-links/show/self`
**Group Route:** `/user-links`
**Handler:** `ShowUserLinksSelf`
**Authentication:** None

**Response Format**: JSON

#### `DELETE /user-links/delete/:id`
**Group Route:** `/user-links`
**Handler:** `DeleteLinks`
**Authentication:** Required (JWT)

**Response Format**: JSON

### 📁 Guest_books Routes
Kumpulan endpoint untuk modul **Guest_books**.

#### `POST /guest-book/create`
**Group Route:** `/guest-book`
**Handler:** `CreateGuest`
**Authentication:** Required (JWT)

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
**Group Route:** `/guest-book`
**Handler:** `EditGuest`
**Authentication:** Required (JWT)

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
**Group Route:** `/guest-book`
**Handler:** `GetAllGuestBook`
**Authentication:** None

**Response Format**: JSON

#### `GET /guest-book/show/:id`
**Group Route:** `/guest-book`
**Handler:** `GetGuestBook`
**Authentication:** None

**Response Format**: JSON

#### `PUT /guest-book/approve/:id`
**Group Route:** `/guest-book`
**Handler:** `ApprovedGuest`
**Authentication:** None

**Response Format**: JSON

#### `PUT /guest-book/reject/:id`
**Group Route:** `/guest-book`
**Handler:** `RejectGuestBook`
**Authentication:** None

**Response Format**: JSON

### 📁 Forum_post Routes
Kumpulan endpoint untuk modul **Forum_post**.

#### `POST /forum-post/create`
**Group Route:** `/forum-post`
**Handler:** `CreatePost`
**Authentication:** Required (JWT)

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
**Group Route:** `/forum-post`
**Handler:** `EditPost`
**Authentication:** Required (JWT)

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
**Group Route:** `/forum-post`
**Handler:** `DeletePost`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `GET /forum-post/showAll`
**Group Route:** `/forum-post`
**Handler:** `GetAllPost`
**Authentication:** None

**Response Format**: JSON

#### `GET /forum-post/show/:id`
**Group Route:** `/forum-post`
**Handler:** `ShowPostById`
**Authentication:** None

**Response Format**: JSON

#### `PUT /forum-post/changeStatus`
**Group Route:** `/forum-post`
**Handler:** `ChangePostStatus`
**Authentication:** None

**Response Format**: JSON

#### `PUT /forum-post/upvote/:id`
**Group Route:** `/forum-post`
**Handler:** `UpvotePost`
**Authentication:** None

**Response Format**: JSON

#### `PUT /forum-post/downvote/:id`
**Group Route:** `/forum-post`
**Handler:** `DownvotePost`
**Authentication:** None

**Response Format**: JSON

#### `POST /forum-post/reply/:id`
**Group Route:** `/forum-post`
**Handler:** `CreateReply`
**Authentication:** None

**Response Format**: JSON

#### `GET /forum-post/post-replies/:id`
**Group Route:** `/forum-post`
**Handler:** `GetPostWithReplies`
**Authentication:** None

**Response Format**: JSON

#### `GET /forum-post/replies/:id`
**Group Route:** `/forum-post`
**Handler:** `GetNestedReplies`
**Authentication:** None

**Response Format**: JSON

### 📁 Portal Routes
Kumpulan endpoint untuk modul **Portal**.

#### `POST /portal/create`
**Group Route:** `/portal`
**Handler:** `CreatePortal`
**Authentication:** Required (JWT)

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
**Group Route:** `/portal`
**Handler:** `EditPortal`
**Authentication:** Required (JWT)

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
**Group Route:** `/portal`
**Handler:** `ShowPortal`
**Authentication:** None

**Response Format**: JSON

#### `GET /portal/showAll`
**Group Route:** `/portal`
**Handler:** `ShowAllPortal`
**Authentication:** None

**Response Format**: JSON

#### `DELETE /portal/delete/:id`
**Group Route:** `/portal`
**Handler:** `DeletePortal`
**Authentication:** Required (JWT)

**Response Format**: JSON

### 📁 Eskul Routes
Kumpulan endpoint untuk modul **Eskul**.

#### `POST /eskul/create`
**Group Route:** `/eskul`
**Handler:** `CreateEskul`
**Authentication:** Required (JWT)

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
**Group Route:** `/eskul`
**Handler:** `EditEskul`
**Authentication:** Required (JWT)

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
**Group Route:** `/eskul`
**Handler:** `ShowEskul`
**Authentication:** None

**Response Format**: JSON

#### `GET /eskul/showAll`
**Group Route:** `/eskul`
**Handler:** `ShowAllEskul`
**Authentication:** None

**Response Format**: JSON

#### `DELETE /eskul/delete/:id`
**Group Route:** `/eskul`
**Handler:** `DeleteEskul`
**Authentication:** Required (JWT)

**Response Format**: JSON

### 📁 Image Routes
Kumpulan endpoint untuk modul **Image**.

#### `POST /image/add`
**Group Route:** `/image`
**Handler:** `AddImage`
**Authentication:** None

**Response Format**: JSON

#### `DELETE /image/delete/:id`
**Group Route:** `/image`
**Handler:** `DeleteImage`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `GET /image/showAll`
**Group Route:** `/image`
**Handler:** `ShowAllImages`
**Authentication:** None

**Response Format**: JSON

#### `GET /image/show`
**Group Route:** `/image`
**Handler:** `ShowImagesByCategory`
**Authentication:** None

**Response Format**: JSON

### 📁 Mading Routes
Kumpulan endpoint untuk modul **Mading**.

#### `POST /mading/create`
**Group Route:** `/mading`
**Handler:** `CreateMading`
**Authentication:** Required (JWT)

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
**Group Route:** `/mading`
**Handler:** `EditMading`
**Authentication:** Required (JWT)

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
**Group Route:** `/mading`
**Handler:** `ShowMading`
**Authentication:** None

**Response Format**: JSON

#### `GET /mading/showAll`
**Group Route:** `/mading`
**Handler:** `ShowAllMadings`
**Authentication:** None

**Response Format**: JSON

#### `DELETE /mading/delete/:id`
**Group Route:** `/mading`
**Handler:** `DeleteMading`
**Authentication:** Required (JWT)

**Response Format**: JSON

#### `PUT /madingchange-status/:id`
**Group Route:** `/mading`
**Handler:** `ChangeStatusMading`
**Authentication:** None

**Response Format**: JSON
