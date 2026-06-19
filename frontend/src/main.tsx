import { AddToCart } from './islands/AddToCart'
import { mountIsland } from './mount'

function mountIfPresent<P extends object>(
  elementId: string,
  Component: React.ComponentType<P>,
  getProps: (el: HTMLElement) => P
) {
  const el = document.getElementById(elementId)
  if (el) {
    mountIsland({ el, Component, props: getProps(el) })
  }
}

mountIfPresent('product-overview', AddToCart, el => ({
  productId: el.dataset.productId!,
}))
