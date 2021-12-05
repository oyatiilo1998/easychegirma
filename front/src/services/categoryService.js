import reqGenerator from "./reqGenerator"

const categoryService = {
  getCategoryList: (params) => reqGenerator.get('/category', { params }),
  getCategoryById: (id) => reqGenerator.get(`/category/${id}`),
}

export default categoryService