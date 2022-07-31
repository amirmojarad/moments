// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"moments/ent/migrate"

	"moments/ent/channel"
	"moments/ent/channelpost"
	"moments/ent/file"
	"moments/ent/group"
	"moments/ent/post"
	"moments/ent/privatechat"
	"moments/ent/publicchat"
	"moments/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Channel is the client for interacting with the Channel builders.
	Channel *ChannelClient
	// ChannelPost is the client for interacting with the ChannelPost builders.
	ChannelPost *ChannelPostClient
	// File is the client for interacting with the File builders.
	File *FileClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
	// PrivateChat is the client for interacting with the PrivateChat builders.
	PrivateChat *PrivateChatClient
	// PublicChat is the client for interacting with the PublicChat builders.
	PublicChat *PublicChatClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Channel = NewChannelClient(c.config)
	c.ChannelPost = NewChannelPostClient(c.config)
	c.File = NewFileClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.Post = NewPostClient(c.config)
	c.PrivateChat = NewPrivateChatClient(c.config)
	c.PublicChat = NewPublicChatClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Channel:     NewChannelClient(cfg),
		ChannelPost: NewChannelPostClient(cfg),
		File:        NewFileClient(cfg),
		Group:       NewGroupClient(cfg),
		Post:        NewPostClient(cfg),
		PrivateChat: NewPrivateChatClient(cfg),
		PublicChat:  NewPublicChatClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Channel:     NewChannelClient(cfg),
		ChannelPost: NewChannelPostClient(cfg),
		File:        NewFileClient(cfg),
		Group:       NewGroupClient(cfg),
		Post:        NewPostClient(cfg),
		PrivateChat: NewPrivateChatClient(cfg),
		PublicChat:  NewPublicChatClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Channel.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Channel.Use(hooks...)
	c.ChannelPost.Use(hooks...)
	c.File.Use(hooks...)
	c.Group.Use(hooks...)
	c.Post.Use(hooks...)
	c.PrivateChat.Use(hooks...)
	c.PublicChat.Use(hooks...)
	c.User.Use(hooks...)
}

// ChannelClient is a client for the Channel schema.
type ChannelClient struct {
	config
}

// NewChannelClient returns a client for the Channel from the given config.
func NewChannelClient(c config) *ChannelClient {
	return &ChannelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `channel.Hooks(f(g(h())))`.
func (c *ChannelClient) Use(hooks ...Hook) {
	c.hooks.Channel = append(c.hooks.Channel, hooks...)
}

// Create returns a builder for creating a Channel entity.
func (c *ChannelClient) Create() *ChannelCreate {
	mutation := newChannelMutation(c.config, OpCreate)
	return &ChannelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Channel entities.
func (c *ChannelClient) CreateBulk(builders ...*ChannelCreate) *ChannelCreateBulk {
	return &ChannelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Channel.
func (c *ChannelClient) Update() *ChannelUpdate {
	mutation := newChannelMutation(c.config, OpUpdate)
	return &ChannelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChannelClient) UpdateOne(ch *Channel) *ChannelUpdateOne {
	mutation := newChannelMutation(c.config, OpUpdateOne, withChannel(ch))
	return &ChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChannelClient) UpdateOneID(id int) *ChannelUpdateOne {
	mutation := newChannelMutation(c.config, OpUpdateOne, withChannelID(id))
	return &ChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Channel.
func (c *ChannelClient) Delete() *ChannelDelete {
	mutation := newChannelMutation(c.config, OpDelete)
	return &ChannelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChannelClient) DeleteOne(ch *Channel) *ChannelDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ChannelClient) DeleteOneID(id int) *ChannelDeleteOne {
	builder := c.Delete().Where(channel.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChannelDeleteOne{builder}
}

// Query returns a query builder for Channel.
func (c *ChannelClient) Query() *ChannelQuery {
	return &ChannelQuery{
		config: c.config,
	}
}

// Get returns a Channel entity by its id.
func (c *ChannelClient) Get(ctx context.Context, id int) (*Channel, error) {
	return c.Query().Where(channel.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChannelClient) GetX(ctx context.Context, id int) *Channel {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ChannelClient) Hooks() []Hook {
	return c.hooks.Channel
}

// ChannelPostClient is a client for the ChannelPost schema.
type ChannelPostClient struct {
	config
}

// NewChannelPostClient returns a client for the ChannelPost from the given config.
func NewChannelPostClient(c config) *ChannelPostClient {
	return &ChannelPostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `channelpost.Hooks(f(g(h())))`.
func (c *ChannelPostClient) Use(hooks ...Hook) {
	c.hooks.ChannelPost = append(c.hooks.ChannelPost, hooks...)
}

// Create returns a builder for creating a ChannelPost entity.
func (c *ChannelPostClient) Create() *ChannelPostCreate {
	mutation := newChannelPostMutation(c.config, OpCreate)
	return &ChannelPostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChannelPost entities.
func (c *ChannelPostClient) CreateBulk(builders ...*ChannelPostCreate) *ChannelPostCreateBulk {
	return &ChannelPostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChannelPost.
func (c *ChannelPostClient) Update() *ChannelPostUpdate {
	mutation := newChannelPostMutation(c.config, OpUpdate)
	return &ChannelPostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChannelPostClient) UpdateOne(cp *ChannelPost) *ChannelPostUpdateOne {
	mutation := newChannelPostMutation(c.config, OpUpdateOne, withChannelPost(cp))
	return &ChannelPostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChannelPostClient) UpdateOneID(id int) *ChannelPostUpdateOne {
	mutation := newChannelPostMutation(c.config, OpUpdateOne, withChannelPostID(id))
	return &ChannelPostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChannelPost.
func (c *ChannelPostClient) Delete() *ChannelPostDelete {
	mutation := newChannelPostMutation(c.config, OpDelete)
	return &ChannelPostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChannelPostClient) DeleteOne(cp *ChannelPost) *ChannelPostDeleteOne {
	return c.DeleteOneID(cp.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ChannelPostClient) DeleteOneID(id int) *ChannelPostDeleteOne {
	builder := c.Delete().Where(channelpost.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChannelPostDeleteOne{builder}
}

// Query returns a query builder for ChannelPost.
func (c *ChannelPostClient) Query() *ChannelPostQuery {
	return &ChannelPostQuery{
		config: c.config,
	}
}

// Get returns a ChannelPost entity by its id.
func (c *ChannelPostClient) Get(ctx context.Context, id int) (*ChannelPost, error) {
	return c.Query().Where(channelpost.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChannelPostClient) GetX(ctx context.Context, id int) *ChannelPost {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ChannelPostClient) Hooks() []Hook {
	return c.hooks.ChannelPost
}

// FileClient is a client for the File schema.
type FileClient struct {
	config
}

// NewFileClient returns a client for the File from the given config.
func NewFileClient(c config) *FileClient {
	return &FileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `file.Hooks(f(g(h())))`.
func (c *FileClient) Use(hooks ...Hook) {
	c.hooks.File = append(c.hooks.File, hooks...)
}

// Create returns a builder for creating a File entity.
func (c *FileClient) Create() *FileCreate {
	mutation := newFileMutation(c.config, OpCreate)
	return &FileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of File entities.
func (c *FileClient) CreateBulk(builders ...*FileCreate) *FileCreateBulk {
	return &FileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for File.
func (c *FileClient) Update() *FileUpdate {
	mutation := newFileMutation(c.config, OpUpdate)
	return &FileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FileClient) UpdateOne(f *File) *FileUpdateOne {
	mutation := newFileMutation(c.config, OpUpdateOne, withFile(f))
	return &FileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FileClient) UpdateOneID(id int) *FileUpdateOne {
	mutation := newFileMutation(c.config, OpUpdateOne, withFileID(id))
	return &FileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for File.
func (c *FileClient) Delete() *FileDelete {
	mutation := newFileMutation(c.config, OpDelete)
	return &FileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FileClient) DeleteOne(f *File) *FileDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *FileClient) DeleteOneID(id int) *FileDeleteOne {
	builder := c.Delete().Where(file.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FileDeleteOne{builder}
}

// Query returns a query builder for File.
func (c *FileClient) Query() *FileQuery {
	return &FileQuery{
		config: c.config,
	}
}

// Get returns a File entity by its id.
func (c *FileClient) Get(ctx context.Context, id int) (*File, error) {
	return c.Query().Where(file.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FileClient) GetX(ctx context.Context, id int) *File {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FileClient) Hooks() []Hook {
	return c.hooks.File
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a builder for creating a Group entity.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// PostClient is a client for the Post schema.
type PostClient struct {
	config
}

// NewPostClient returns a client for the Post from the given config.
func NewPostClient(c config) *PostClient {
	return &PostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `post.Hooks(f(g(h())))`.
func (c *PostClient) Use(hooks ...Hook) {
	c.hooks.Post = append(c.hooks.Post, hooks...)
}

// Create returns a builder for creating a Post entity.
func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Post entities.
func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
	return &PostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Post.
func (c *PostClient) Update() *PostUpdate {
	mutation := newPostMutation(c.config, OpUpdate)
	return &PostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PostClient) UpdateOne(po *Post) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPost(po))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PostClient) UpdateOneID(id int) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPostID(id))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Post.
func (c *PostClient) Delete() *PostDelete {
	mutation := newPostMutation(c.config, OpDelete)
	return &PostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PostClient) DeleteOne(po *Post) *PostDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PostClient) DeleteOneID(id int) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

// Query returns a query builder for Post.
func (c *PostClient) Query() *PostQuery {
	return &PostQuery{
		config: c.config,
	}
}

// Get returns a Post entity by its id.
func (c *PostClient) Get(ctx context.Context, id int) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PostClient) GetX(ctx context.Context, id int) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Post.
func (c *PostClient) QueryOwner(po *Post) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.OwnerTable, post.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PostClient) Hooks() []Hook {
	return c.hooks.Post
}

// PrivateChatClient is a client for the PrivateChat schema.
type PrivateChatClient struct {
	config
}

// NewPrivateChatClient returns a client for the PrivateChat from the given config.
func NewPrivateChatClient(c config) *PrivateChatClient {
	return &PrivateChatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `privatechat.Hooks(f(g(h())))`.
func (c *PrivateChatClient) Use(hooks ...Hook) {
	c.hooks.PrivateChat = append(c.hooks.PrivateChat, hooks...)
}

// Create returns a builder for creating a PrivateChat entity.
func (c *PrivateChatClient) Create() *PrivateChatCreate {
	mutation := newPrivateChatMutation(c.config, OpCreate)
	return &PrivateChatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PrivateChat entities.
func (c *PrivateChatClient) CreateBulk(builders ...*PrivateChatCreate) *PrivateChatCreateBulk {
	return &PrivateChatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PrivateChat.
func (c *PrivateChatClient) Update() *PrivateChatUpdate {
	mutation := newPrivateChatMutation(c.config, OpUpdate)
	return &PrivateChatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PrivateChatClient) UpdateOne(pc *PrivateChat) *PrivateChatUpdateOne {
	mutation := newPrivateChatMutation(c.config, OpUpdateOne, withPrivateChat(pc))
	return &PrivateChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PrivateChatClient) UpdateOneID(id int) *PrivateChatUpdateOne {
	mutation := newPrivateChatMutation(c.config, OpUpdateOne, withPrivateChatID(id))
	return &PrivateChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PrivateChat.
func (c *PrivateChatClient) Delete() *PrivateChatDelete {
	mutation := newPrivateChatMutation(c.config, OpDelete)
	return &PrivateChatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PrivateChatClient) DeleteOne(pc *PrivateChat) *PrivateChatDeleteOne {
	return c.DeleteOneID(pc.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PrivateChatClient) DeleteOneID(id int) *PrivateChatDeleteOne {
	builder := c.Delete().Where(privatechat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PrivateChatDeleteOne{builder}
}

// Query returns a query builder for PrivateChat.
func (c *PrivateChatClient) Query() *PrivateChatQuery {
	return &PrivateChatQuery{
		config: c.config,
	}
}

// Get returns a PrivateChat entity by its id.
func (c *PrivateChatClient) Get(ctx context.Context, id int) (*PrivateChat, error) {
	return c.Query().Where(privatechat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PrivateChatClient) GetX(ctx context.Context, id int) *PrivateChat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFirstUser queries the first_user edge of a PrivateChat.
func (c *PrivateChatClient) QueryFirstUser(pc *PrivateChat) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(privatechat.Table, privatechat.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, privatechat.FirstUserTable, privatechat.FirstUserColumn),
		)
		fromV = sqlgraph.Neighbors(pc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySecondUser queries the second_user edge of a PrivateChat.
func (c *PrivateChatClient) QuerySecondUser(pc *PrivateChat) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(privatechat.Table, privatechat.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, privatechat.SecondUserTable, privatechat.SecondUserColumn),
		)
		fromV = sqlgraph.Neighbors(pc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PrivateChatClient) Hooks() []Hook {
	return c.hooks.PrivateChat
}

// PublicChatClient is a client for the PublicChat schema.
type PublicChatClient struct {
	config
}

// NewPublicChatClient returns a client for the PublicChat from the given config.
func NewPublicChatClient(c config) *PublicChatClient {
	return &PublicChatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `publicchat.Hooks(f(g(h())))`.
func (c *PublicChatClient) Use(hooks ...Hook) {
	c.hooks.PublicChat = append(c.hooks.PublicChat, hooks...)
}

// Create returns a builder for creating a PublicChat entity.
func (c *PublicChatClient) Create() *PublicChatCreate {
	mutation := newPublicChatMutation(c.config, OpCreate)
	return &PublicChatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PublicChat entities.
func (c *PublicChatClient) CreateBulk(builders ...*PublicChatCreate) *PublicChatCreateBulk {
	return &PublicChatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PublicChat.
func (c *PublicChatClient) Update() *PublicChatUpdate {
	mutation := newPublicChatMutation(c.config, OpUpdate)
	return &PublicChatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PublicChatClient) UpdateOne(pc *PublicChat) *PublicChatUpdateOne {
	mutation := newPublicChatMutation(c.config, OpUpdateOne, withPublicChat(pc))
	return &PublicChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PublicChatClient) UpdateOneID(id int) *PublicChatUpdateOne {
	mutation := newPublicChatMutation(c.config, OpUpdateOne, withPublicChatID(id))
	return &PublicChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PublicChat.
func (c *PublicChatClient) Delete() *PublicChatDelete {
	mutation := newPublicChatMutation(c.config, OpDelete)
	return &PublicChatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PublicChatClient) DeleteOne(pc *PublicChat) *PublicChatDeleteOne {
	return c.DeleteOneID(pc.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PublicChatClient) DeleteOneID(id int) *PublicChatDeleteOne {
	builder := c.Delete().Where(publicchat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PublicChatDeleteOne{builder}
}

// Query returns a query builder for PublicChat.
func (c *PublicChatClient) Query() *PublicChatQuery {
	return &PublicChatQuery{
		config: c.config,
	}
}

// Get returns a PublicChat entity by its id.
func (c *PublicChatClient) Get(ctx context.Context, id int) (*PublicChat, error) {
	return c.Query().Where(publicchat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PublicChatClient) GetX(ctx context.Context, id int) *PublicChat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PublicChatClient) Hooks() []Hook {
	return c.hooks.PublicChat
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPosts queries the posts edge of a User.
func (c *UserClient) QueryPosts(u *User) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PostsTable, user.PostsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFollowers queries the followers edge of a User.
func (c *UserClient) QueryFollowers(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.FollowersTable, user.FollowersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFollowing queries the following edge of a User.
func (c *UserClient) QueryFollowing(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.FollowingTable, user.FollowingPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMyPvChats queries the my_pv_chats edge of a User.
func (c *UserClient) QueryMyPvChats(u *User) *PrivateChatQuery {
	query := &PrivateChatQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(privatechat.Table, privatechat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.MyPvChatsTable, user.MyPvChatsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOtherPvChats queries the other_pv_chats edge of a User.
func (c *UserClient) QueryOtherPvChats(u *User) *PrivateChatQuery {
	query := &PrivateChatQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(privatechat.Table, privatechat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.OtherPvChatsTable, user.OtherPvChatsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
