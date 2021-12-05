import reqGenerator from "./reqGenerator"

const productService = {
  getProductList: (params) => reqGenerator.get('/product', { params }),
  getProductById: (id) => reqGenerator.get(`/product/${id}`),
}

export default productService