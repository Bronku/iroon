function addCake(id, name, price, amount) {
  const cake = document.querySelector(`[name="cake[${id}]"]`);
  if (cake != null) {
    cake.value = Number(cake.value) + 1;
    return;
  }
  let template = document.getElementById("basket_element_template");
  let tr = template.content.querySelector("tr").cloneNode(true);
  tr.querySelector(".cake-name").innerText = name;
  tr.querySelector(".cake-price").innerText = `${price}PLN`;
  tr.querySelector(".cake-total").innerText = `${price * amount}PLN`;
  tr.querySelector("button").onclick = () => {
    tr.remove();
  };
  let input = tr.querySelector("input");
  input.onchange = () => {
    tr.querySelector(".cake-total").innerText = `${price * input.value}PLN`;
  };
  input.value = amount;
  input.name = `cake[${id}]`;
  document.querySelector("#basket_table").appendChild(tr);
}
function toggleDetails(row) {
  const detailsRow = row.nextElementSibling;
  if (detailsRow.classList.contains("hidden")) {
    detailsRow.classList.remove("hidden");
    return;
  }
  detailsRow.classList.add("hidden");
}
