import ProductCard from "../../../components/ProductCard"
import "./style.scss"

const ProductsBlock = ({openModal, productsList, hiddenTitle, setSelectedProduct}) => {
  return (
    <>
      {!hiddenTitle && <h2 className="main-title">Список скидок</h2>}
      <div className="ProductsBlock">

        {
          productsList?.map((product, index) => (
            <ProductCard className={`${(index === 4 || index === 9) ? 'first' : '' }`} openModal={openModal} product={product} setSelectedProduct={setSelectedProduct}  />
          ))
        }
      
      </div>
    </>
  )
}

export default ProductsBlock
