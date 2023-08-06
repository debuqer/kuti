Here is a draft RFC and software spec for a basic blog application:

## Blog Application RFC

### Overview

This RFC proposes the development of a blog application that allows user to create posts, categorize them, and share them publicly. User should be able to create/edit/delete posts, add tags/categories, and view all published posts.

### Posts

- A post consists of a title, content body, publication date, tags
- User can create new text-based posts 
- User can edit and delete posts they have created
- Post can be assigned descriptive tags by the author
- Post have a published/draft status
- Only published posts are publicly visible

### Tags

- Users can create and assign tags to posts 
- Tags help categorize posts 
- Users can click on a tag to view all posts with that tag

### User Interface

- Simple, clean UI with a homepage, post creation form, and post lists
- Homepage displays a list of published blog post titles/summaries 
- Clicking a post opens the full post content
- Navigation bar allows accessing posts by tag

### Permissions

- Public users can view published posts
- User can create/edit/delete their own posts
- User can edit their account settings
- User can edit all posts

### Backend API Endpoints

- `/api/posts` - Get all posts, Get post by id, Create post, Update post, Delete post
- `/api/tags` - Get all tags, Create tag, Update tag, Delete tag

### Data Schema

```
Author
  - name
  - username 
  - email
  
Post
  - id
  - title
  - content
  - publication_date
  - status (draft, published)
  - tags
  
Tag
  - id 
  - name

```