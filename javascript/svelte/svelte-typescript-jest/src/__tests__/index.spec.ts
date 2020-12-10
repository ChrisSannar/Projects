// import '@testing-library/jest-dom/extend-expect'
import { render } from '@testing-library/svelte'

import index from "../index.svelte"

describe("testing", () => {
  it("works", () => {
    const { getByText } = render(index);

    expect(getByText('Svelte')).toHaveTextContent()

    // let x = 2;
    // expect(x + x).toEqual(4);
  })
})