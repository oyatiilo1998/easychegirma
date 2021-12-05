import Button from "../Button"
import "./style.scss"

const ProductCard = ({ className, openModal, product, setSelectedProduct }) => {

  const clickHandler = () => {
    setSelectedProduct(product)
    openModal()
  }

  return (
    <div className={`ProductCard ${className}`} onClick={clickHandler} >
      <div className="top-side">
        <img
          src={product?.image_url}
          alt=""
          className="product-image"
        />
        <div className="product-name">
          {product?.name}
        </div>
        <div className="category-name" style={{ color: 'silver', margin: '5px 0' }} >{product?.owner?.name}</div>
        <div className="category-name">{product?.category?.name}</div>
      </div>

      <div className="bottom">
        <div className="discount">Скидка {product?.discount_amount}%</div>
        <div className="price-row">
          <span className="old-price">{product?.price_before} UZS</span> -{" "}
          <span className="new-price">{product?.price_after} UZS</span>
        </div>

        <div className="btn-block">
          <a href={`${product?.url}`} target="_blank" rel="noreferrer"  onClick={e => e.stopPropagation()} ><Button className="btn primary">Перейти в магазин</Button></a>
        </div>

      </div>
    </div>
  )
}

export default ProductCard
