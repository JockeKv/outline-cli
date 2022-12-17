package api

import "time"

// Attachment defines model for Attachment.
type Attachment struct {
	ContentType *string `json:"contentType,omitempty"`

	// DocumentId Identifier for the associated document, if any.
	DocumentId *string  `json:"documentId,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Size       *float32 `json:"size,omitempty"`
	Url        *string  `json:"url,omitempty"`
}

// Auth defines model for Auth.
type Auth struct {
	Team *Team `json:"team,omitempty"`
	User *User `json:"user,omitempty"`
}

// Collection defines model for Collection.
type Collection struct {
	// Color A color representing the collection, this is used to help make collections more identifiable in the UI. It should be in HEX format including the #
	Color *string `json:"color,omitempty"`

	// CreatedAt The date and time that this object was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt The date and time that this object was deleted
	DeletedAt *time.Time `json:"deletedAt"`

	// Description A description of the collection, may contain markdown formatting
	Description *string           `json:"description,omitempty"`
	Documents   *[]NavigationNode `json:"documents,omitempty"`

	// Icon A string that represents an icon in the outline-icons package
	Icon *string `json:"icon,omitempty"`

	// Id Unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Index The position of the collection in the sidebar
	Index *string `json:"index,omitempty"`

	// Name The name of the collection.
	Name       *string `json:"name,omitempty"`
	Permission *string `json:"permission,omitempty"`

	// Sort The sort of documents in the collection. Note that not all API responses respect this and it is left as a frontend concern to implement.
	Sort *struct {
		Direction *string `json:"direction,omitempty"`
		Field     *string `json:"field,omitempty"`
	} `json:"sort,omitempty"`

	// UpdatedAt The date and time that this object was last changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// CollectionGroupMembership defines model for CollectionGroupMembership.
type CollectionGroupMembership struct {
	// CollectionId Identifier for the associated collection.
	CollectionId *string `json:"collectionId,omitempty"`

	// GroupId Identifier for the associated group.
	GroupId *string `json:"groupId,omitempty"`

	// Id Unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Permission The permissions that this membership grants the users in the group
	Permission *string `json:"permission,omitempty"`
}

// Document defines model for Document.
type Document struct {
	// ArchivedAt The date and time that this object was archived
	ArchivedAt    *time.Time `json:"archivedAt,omitempty"`
	Collaborators *[]User    `json:"collaborators,omitempty"`

	// CollectionId Identifier for the associated collection.
	CollectionId *string `json:"collectionId,omitempty"`

	// CreatedAt The date and time that this object was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	CreatedBy *User      `json:"createdBy,omitempty"`

	// DeletedAt The date and time that this object was deleted
	DeletedAt *time.Time `json:"deletedAt"`

	// Emoji An emoji associated with the document.
	Emoji *string `json:"emoji,omitempty"`

	// FullWidth Whether this document should be displayed in a full-width view.
	FullWidth *bool `json:"fullWidth,omitempty"`

	// Id Unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// ParentDocumentId Identifier for the document this is a child of, if any.
	ParentDocumentId *string `json:"parentDocumentId,omitempty"`

	// Pinned Whether this document is pinned in the collection
	Pinned *bool `json:"pinned,omitempty"`

	// PublishedAt The date and time that this object was published
	PublishedAt *time.Time `json:"publishedAt"`

	// Revision A number that is auto incrementing with every revision of the document that is saved
	Revision *float32 `json:"revision,omitempty"`

	// Template Whether this document is a template
	Template *bool `json:"template,omitempty"`

	// TemplateId Unique identifier for the template this document was created from, if any
	TemplateId *string `json:"templateId,omitempty"`

	// Text The text content of the document, contains markdown formatting
	Text *string `json:"text,omitempty"`

	// Title The title of the document.
	Title *string `json:"title,omitempty"`

	// UpdatedAt The date and time that this object was last changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	UpdatedBy *User      `json:"updatedBy,omitempty"`

	// UrlId A short unique ID that can be used to identify the document as an alternative to the UUID
	UrlId *string `json:"urlId,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Error   string `json:"error,omitempty"`
	Ok      bool   `json:"ok,omitempty"`
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// Group defines model for Group.
type Group struct {
	// CreatedAt The date and time that this object was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Id Unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// MemberCount The number of users that are members of the group
	MemberCount *float32 `json:"memberCount,omitempty"`

	// Name The name of this group.
	Name *string `json:"name,omitempty"`

	// UpdatedAt The date and time that this object was last changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// GroupMembership defines model for GroupMembership.
type GroupMembership struct {
	// GroupId Identifier for the associated group.
	GroupId *string `json:"groupId,omitempty"`

	// Id Unique identifier for the object.
	Id   *string `json:"id,omitempty"`
	User *User   `json:"user,omitempty"`

	// UserId Identifier for the associated user.
	UserId *string `json:"userId,omitempty"`
}

// Invite defines model for Invite.
type Invite struct {
	// Email The email address to invite
	Email *string `json:"email,omitempty"`

	// Name The full name of the user being invited
	Name *string `json:"name,omitempty"`
	Role *string `json:"role,omitempty"`
}

// NavigationNode defines model for NavigationNode.
type NavigationNode struct {
	Children *[]NavigationNode `json:"children,omitempty"`

	// Id Unique identifier for the document.
	Id    *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Url   *string `json:"url,omitempty"`
}

// Pagination defines model for Pagination.
type Pagination struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// Revision defines model for Revision.
type Revision struct {
	// CreatedAt Date and time when this revision was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	CreatedBy *User      `json:"createdBy,omitempty"`

	// DocumentId Identifier for the associated document.
	DocumentId *string `json:"documentId,omitempty"`

	// Id Unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Text Body of the document, may contain markdown formatting
	Text *string `json:"text,omitempty"`

	// Title Title of the document.
	Title *string `json:"title,omitempty"`
}

// Share defines model for Share.
type Share struct {
	// CreatedAt Date and time when this share was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	CreatedBy *User      `json:"createdBy,omitempty"`

	// DocumentTitle Title of the shared document.
	DocumentTitle *string `json:"documentTitle,omitempty"`

	// DocumentUrl URL of the original document.
	DocumentUrl *string `json:"documentUrl,omitempty"`

	// Id Unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// IncludeChildDocuments If to also give permission to view documents nested beneath this one.
	IncludeChildDocuments *bool `json:"includeChildDocuments,omitempty"`

	// LastAccessedAt Date and time when this share was last viewed
	LastAccessedAt *time.Time `json:"lastAccessedAt,omitempty"`

	// Published If true the share can be loaded without a user account.
	Published *bool `json:"published,omitempty"`

	// UpdatedAt Date and time when this share was edited
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Url URL of the publicly shared document.
	Url *string `json:"url,omitempty"`
}

// Sorting defines model for Sorting.
type Sorting struct {
	Direction *string `json:"direction,omitempty"`
	Sort      *string `json:"sort,omitempty"`
}

// Team defines model for Team.
type Team struct {
	AllowedDomains *[]string `json:"allowedDomains,omitempty"`

	// AvatarUrl The URL for the image associated with this team, it will be displayed in the team switcher and in the top left of the knowledge base along with the name.
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// CollaborativeEditing Whether this team has collaborative editing in documents globally enabled.
	CollaborativeEditing *bool `json:"collaborativeEditing,omitempty"`

	// DefaultCollectionId If set then the referenced collection is where users will be redirected to after signing in instead of the Home screen
	DefaultCollectionId *string `json:"defaultCollectionId,omitempty"`

	// DefaultUserRole If set then this role will be used as the default for users that signup via SSO
	DefaultUserRole *string `json:"defaultUserRole,omitempty"`

	// DocumentEmbeds Whether this team has embeds in documents globally enabled. It can be disabled to reduce potential data leakage to third parties.
	DocumentEmbeds *bool `json:"documentEmbeds,omitempty"`

	// GuestSignin Whether this team has guest signin enabled. Guests can signin with an email address and are not required to have a Google Workspace/Slack SSO account once invited.
	GuestSignin *bool `json:"guestSignin,omitempty"`

	// Id Unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// InviteRequired Whether an invite is required to join this team, if false users may join with a linked SSO provider.
	InviteRequired *bool `json:"inviteRequired,omitempty"`

	// MemberCollectionCreate Whether members are allowed to create new collections. If false then only admins can create collections.
	MemberCollectionCreate *bool `json:"memberCollectionCreate,omitempty"`

	// Name The name of this team, it is usually auto-generated when the first SSO connection is made but can be changed if neccessary.
	Name *string `json:"name,omitempty"`

	// Sharing Whether this team has share links globally enabled. If this value is false then all sharing UI and APIs are disabled.
	Sharing *bool `json:"sharing,omitempty"`

	// Subdomain Represents the subdomain at which this team's knowledge base can be accessed.
	Subdomain *string `json:"subdomain,omitempty"`

	// Url The fully qualified URL at which this team's knowledge base can be accessed.
	Url *string `json:"url,omitempty"`
}

// User defines model for User.
type User struct {
	// AvatarUrl The URL for the image associated with this user, it will be displayed in the application UI and email notifications.
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// CreatedAt The date and time that this user first signed in or was invited as a guest.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Email The email associated with this user, it is migrated from Slack or Google Workspace when the SSO connection is made but can be changed if neccessary.
	Email *string `json:"email,omitempty"`

	// Id Unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// IsAdmin Whether this user has admin permissions.
	IsAdmin *bool `json:"isAdmin,omitempty"`

	// IsSuspended Whether this user has been suspended.
	IsSuspended *bool `json:"isSuspended,omitempty"`

	// LastActiveAt The last time this user made an API request, this value is updated at most every 5 minutes.
	LastActiveAt *string `json:"lastActiveAt,omitempty"`

	// Name The name of this user, it is migrated from Slack or Google Workspace when the SSO connection is made but can be changed if neccessary.
	Name *string `json:"name,omitempty"`
}
