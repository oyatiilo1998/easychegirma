import { useEffect, useState } from "react"
import { useParams } from "react-router"
import ProductModal from "../../components/ProductModal"
import categoryService from "../../services/categoryService"
import productService from "../../services/productService"
import ProductsBlock from "../main/productsBlock"


const CategoryPage = () => {
  const params = useParams()
  const [modalIsOpen, setModalIsOpen] = useState(false)
  const [productsList, setProductsList] = useState([])
  const [categoryData, setCategoryData] = useState(null)

  const openModal = () => setModalIsOpen(true)
  const closeModal = () => setModalIsOpen(false)

  const fetchCategoryData = () => {
    categoryService
      .getCategoryById(params.id)
      .then(res => setCategoryData(res))
  }

  const fetchProductsList = () => {
    productService
      .getProductList({category_id: params.id})
      .then(res => setProductsList(res?.products))
  }

  useEffect(() => {
    fetchCategoryData()
    fetchProductsList()
  }, [])

  return (
    <div className="container" style={{paddingTop: '20px'}} >
      <h1 className="main-title">{categoryData?.name}</h1>
      <ProductModal isOpen={modalIsOpen} onRequestClose={closeModal} />
      <ProductsBlock openModal={openModal} productsList={productsList} hiddenTitle  />
    </div>
  )
}

export default CategoryPage
