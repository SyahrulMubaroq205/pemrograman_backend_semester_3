import { index, store, update, destroy } from "./FruitController.js";

const main = () => {
  console.log("Initial fruits:", index());
  console.log("After adding Lemon:", store("Lemon"));
  console.log("After updating position 1:", update(1, "Blueberry"));
  console.log("After destroying position 2:", destroy(2));
};

main();
