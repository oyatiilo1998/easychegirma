import reqGenerator from "../utils/reqGenerator";

const userService = {
  getUsers: (params) => reqGenerator.get('/user', { params }),
  getUserById: (id) => reqGenerator.get(`/user/${id}`),
  createUser: (data) => reqGenerator.post('/user', data),
  updateUser: (id, data) => reqGenerator.put(`/user/${id}`, data),
}

export default userService