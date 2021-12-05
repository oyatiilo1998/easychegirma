import reqGenerator from "../utils/reqGenerator";

const categoryService = {
  getCategoryList: (params) => reqGenerator.get('/category', { params }),
  getCategoryById: (id) => reqGenerator.get(`/category/${id}`),
  createCategory: (data) => reqGenerator.post('/category', data),
  updateCategory: (id, data) => reqGenerator.put(`/category/${id}`, data),
}

export default categoryService