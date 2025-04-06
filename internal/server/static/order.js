function addCake(id, name, amount) {
  let template = document.getElementById("basket_element_template");
  let tr = template.content.cloneNode(true);
  console.log(tr);
  const cake = document.getElementById("cake[" + id + "]");
  if (cake != null) {
    const value = Number(cake.getAttribute("value"));
    cake.setAttribute("value", value + 1);
    return;
  }
  const label = document.createElement("label");
  label.innerText = name;
  const input = document.createElement("input");
  input.setAttribute("id", "cake[" + id + "]");
  input.setAttribute("type", "number");
  input.setAttribute("min", "0");
  input.setAttribute("name", "cake[" + id + "]");
  input.setAttribute("value", amount);
  const button = document.createElement("button");
  button.setAttribute("type", "button");
  button.setAttribute("onClick", "removeCake(" + id + ")");
  button.innerText = "usuń";
  const li = document.createElement("li");
  li.appendChild(label);
  li.appendChild(input);
  li.appendChild(button);
  li.setAttribute("id", "li[" + id + "]");
  const ul = document.getElementById("basket");
  ul.appendChild(li);
}
function removeCake(id) {
  document.getElementById("li[" + id + "]").remove();
}
