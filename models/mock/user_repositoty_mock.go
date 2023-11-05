package mock_models

import "chat/models"

type MockUserRepository struct {
	users     map[int]models.User
	chatrooms map[int]models.Chatroom
	currId    int
}

func NewMockUserRepository() *MockUserRepository {
	chats := make(map[int]models.Chatroom)
	chat := new(models.Chatroom)
	chats[1] = *chat
	return &MockUserRepository{
		users:     make(map[int]models.User),
		chatrooms: chats,
	}
}

func (m *MockUserRepository) Fetch(limit int) (response []models.User, err error) {
	if limit == 0 {
		for _, val := range m.users {
			response = append(response, val)
		}
		return
	} else if limit > 0 {
		for i := 1; i <= limit; i++ {
			response = append(response, m.users[i])
		}
		return
	} else {
		err = models.ErrBadParamInput
		return
	}
}

func (m *MockUserRepository) FetchOne(id int) (models.User, error) {
	for key, val := range m.users {
		if key == id {
			return val, nil
		}
	}

	return models.User{}, models.ErrNotFound
}

func (m *MockUserRepository) FetchFewCertain(ids ...int) (response []models.User, err error) {
	for key, val := range m.users {
		for _, id := range ids {
			if id == key {
				response = append(response, val)
			}
		}
	}

	if len(response) == 0 {
		err = models.ErrNotFound
		return
	}

	return
}

func (m *MockUserRepository) Store(user models.User) error {
	m.currId++
	user.ID = m.currId
	m.users[m.currId] = user
	return nil
}

func (m *MockUserRepository) Update(user models.User) error {
	if user.Name == "" || user.Password == "" {
		return models.ErrEmptyFields
	}

	userFound := false
	for i := 0; i < len(m.users); i++ {
		if _, ok := m.users[user.ID]; ok {
			userFound = true
		}
	}

	if userFound {
		m.users[user.ID] = user
		return nil
	} else {
		return models.ErrNotFound
	}
}

func (m *MockUserRepository) Delete(id int) error {
	userFound := false
	for i := 0; i < len(m.users); i++ {
		if _, ok := m.users[id]; ok {
			userFound = true
		}
	}

	if userFound {
		delete(m.users, id)
		return nil
	} else {
		return models.ErrNotFound
	}
}

func (m *MockUserRepository) GetChatters() (chatters []models.User) {
	for _, user := range m.users {
		if user.CurrentChatroomID != 0 {
			chatters = append(chatters, user)
		}
	}

	return
}

func (m *MockUserRepository) GetUserPassword(id int) (string, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user.Password, nil
		}
	}

	return "", models.ErrNotFound
}

func (m *MockUserRepository) GetuserName(id int) (string, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user.Name, nil
		}
	}

	return "", models.ErrNotFound
}

func (m *MockUserRepository) BeforeJoin(uid, cid int) bool {
	userFound := false
	chatroomFound := false
	for _, user := range m.users {
		if user.ID == uid {
			userFound = true
		}
	}

	for _, chat := range m.chatrooms {
		if chat.ID == cid {
			chatroomFound = true
		}
	}

	return userFound && chatroomFound
}
