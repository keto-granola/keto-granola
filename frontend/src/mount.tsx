import type { ComponentType } from 'react'
import { hydrateRoot } from 'react-dom/client'

interface MountOpts<P> {
  el: HTMLElement
  Component: ComponentType<P>
  props: P
}

export function mountIsland<P extends object>({ el, Component, props }: MountOpts<P>) {
  hydrateRoot(el, <Component {...props} />)
}
