import { useState } from 'react'

import { mountIsland } from '../mount'

interface ProductOverviewProps {
  productId: string
}

export function ProductOverview({ productId }: ProductOverviewProps) {
  const [quantity, setQuantity] = useState<number>(1)

  const addToCart = (productId: string, quantity: number) => {
    // TODO: add to cart
    console.warn(productId, quantity)
  }

  return (
    <div>
      <button onClick={() => setQuantity(q => q - 1)}>-</button>
      <span>{quantity}</span>
      <button onClick={() => setQuantity(q => q + 1)}>+</button>
      <button onClick={() => addToCart(productId, quantity)}>Add to cart</button>
    </div>
  )
}

const el = document.getElementById('product-overview')
if (el) {
  const productId = el.dataset.productId
  if (!productId) {
    throw new Error('missing product id')
  }

  mountIsland({
    el,
    Component: ProductOverview,
    props: { productId },
  })
}
