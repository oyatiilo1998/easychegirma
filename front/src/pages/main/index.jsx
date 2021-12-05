import "./style.scss"
import Carousel from "nuka-carousel"
import CatalogsBlock from "./catalogsBlock"
import Footer from "../../components/Footer"
import ProductsBlock from "./productsBlock"
import ProductModal from "../../components/ProductModal"
import { useEffect, useState } from "react"
import categoryService from "../../services/categoryService"
import productService from "../../services/productService"


const MainPage = () => {
  const [modalIsOpen, setModalIsOpen] = useState(false)
  const [categoriesList, setCategoriesList] = useState([])
  const [productsList, setProductsList] = useState([])
  const [categoryLoader, setCategoryLoader] = useState(true)
  const [productLoader, setProductLoader] = useState(true)
  const [selectedProduct, setSelectedProduct] = useState(null)

  const fetchCategoriesList = () => {
    setCategoryLoader(true)
    categoryService
      .getCategoryList()
      .then(res => setCategoriesList(res?.categories))
      .finally(() => setCategoryLoader(false))
  }

  const fetchProductsList = () => {
    setProductLoader(true)
    productService
      .getProductList()
      .then(res => setProductsList(res?.products))
      .finally(() => setProductLoader(false))
  }

  useEffect(() => {
    fetchCategoriesList()
    fetchProductsList()
  }, [])
  

  const openModal = () => setModalIsOpen(true)
  const closeModal = () => setModalIsOpen(false)

  return (
    <div className="MainPage">
      <ProductModal selectedProduct={selectedProduct} isOpen={modalIsOpen} onRequestClose={closeModal} />


      <div className="first-section">
        <div className="content container">
          <div className="carousel-block">
            <Carousel autoplay={true} autoplayInterval={2000} defaultControlsConfig={{nextButtonClassName: 'hidden', prevButtonClassName: 'hidden'}} >
              <img src={'https://texnomart.uz/frontend/web/uploads/slides/2038151920x400_%D1%80%D0%BE%D0%B7%D1%8B%D0%B3%D1%80%D1%8B%D1%88%20%D0%BF%D1%80%D0%B8%D0%B7%D0%BE%D0%B2.png'} alt="img1" />
              <img src={'https://texnomart.uz/frontend/web/uploads/slides/1758561920x400.png'} alt="img3" />
              <img src={'https://texnomart.uz/frontend/web/uploads/slides/3781751920%D1%85400.png'} alt="img2" />
            </Carousel>
          </div>


          <CatalogsBlock categoriesList={categoriesList} />

          <div className="separator" />

          <ProductsBlock openModal={openModal} productsList={productsList} setSelectedProduct={setSelectedProduct} />

        </div>
      </div>
      <Footer />
    </div>
  )
}

export default MainPage
