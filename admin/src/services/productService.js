import reqGenerator from "../utils/reqGenerator";

const productService = {
  getProductList: (params) => reqGenerator.get('/product', { params }),
  getProductById: (id) => reqGenerator.get(`/product/${id}`),
  createProduct: (data) => reqGenerator.post('/product', data),
  updateProduct: (id, data) => reqGenerator.put(`/product/${id}`, data),
  deleteProduct: (id) => reqGenerator.delete(`/product/${id}`)
}

export default productService