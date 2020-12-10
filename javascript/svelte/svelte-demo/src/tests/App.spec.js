import App from "../App.svelte";
import { render } from "@testing-library/svelte";

it("resolves", async () => {});

describe("App Component", () => {
  test("should render component", () => {
    const { container } = render(App);

    expect(container).toContain("Hello");
  });
});
