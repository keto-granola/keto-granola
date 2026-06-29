import type { ComponentType } from 'react'
import { createRoot, hydrateRoot } from 'react-dom/client'

interface MountOpts<P> {
  el: HTMLElement
  Component: ComponentType<P>
  props: P
  hydrate?: boolean
}

export function mountIsland<P extends object>({
  el,
  Component,
  props,
  hydrate = false,
}: MountOpts<P>) {
  if (hydrate) {
    hydrateRoot(el, <Component {...props} />)
  } else {
    createRoot(el).render(<Component {...props} />)
  }
}
