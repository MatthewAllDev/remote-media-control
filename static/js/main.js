const themeLink = document.getElementById("theme-link");
const themeMeta = document.getElementById("theme-meta");
const fullScreenSwitch = document.getElementById("FullScreenSwitch");
const darkModeSwitch = document.getElementById("DarkModeSwitch");

let activePushButton = false,
  keyRepeatTimer;

init();
addListeners();

function init() {
  let currentPage = localStorage.getItem("currentPage");
  if (currentPage) {
    document.getElementById(currentPage + "-tab").classList.add("active");
    document.getElementById(currentPage).classList.add("active", "show");
  } else {
    document.getElementsByClassName("nav-link")[0].classList.add("active");
    document
      .getElementsByClassName("tab-pane")[0]
      .classList.add("active", "show");
  }
  if (Number(localStorage.getItem("useDarkMode"))) {
    applyTheme(true);
  }
  [...document.getElementsByClassName("tab-pane")].forEach((element) => {
    let visibility = localStorage.getItem(element.id + "-visibility");
    if (visibility === null) return;
    showPage(element.id, Number(visibility));
  });
}

function addListeners() {
  const on = (selector, events, listener) => {
    document
      .querySelectorAll(selector)
      .forEach((el) =>
        events
          .split(" ")
          .forEach((event) => el.addEventListener(event, listener)),
      );
  };

  on(".push-button", "pointerdown", pushButtonListenerDown);
  on(".push-button", "pointerup pointercancel", pushButtonListenerUp);
  on(".click-button", "click", clickButtonListener);
  on(".nav-link", "click", (event) => {
    const target = event.currentTarget.getAttribute("data-bs-target");
    if (target) localStorage.setItem("currentPage", target.slice(1));
  });
  on(".page-visibility-switch", "change", pageVisibilitySwitchListener);
  darkModeSwitch.addEventListener("change", darkModeSwitchListener);
  fullScreenSwitch.addEventListener("change", fullScreenSwitchListener);
  document.addEventListener("fullscreenchange", fullScreenChangeListener);
}

function pushButtonListenerDown(event) {
  sendKey(event.currentTarget.attributes.getNamedItem("data-sent-key").value);
  activePushButton = event.currentTarget;
  keyRepeatTimer = setInterval(function () {
    sendKey(activePushButton.attributes.getNamedItem("data-sent-key").value);
  }, 250);
}

function pushButtonListenerUp(event) {
  if (
    event.currentTarget === activePushButton ||
    activePushButton !== false ||
    event.type === "pointerup"
  ) {
    clearInterval(keyRepeatTimer);
    activePushButton = false;
  }
}

function clickButtonListener(event) {
  sendKey(event.currentTarget.attributes.getNamedItem("data-sent-key").value);
}

function darkModeSwitchListener(event) {
  const isDarkMode = event.currentTarget.checked;
  localStorage.setItem("useDarkMode", Number(isDarkMode));
  applyTheme(isDarkMode);
}

function fullScreenSwitchListener(event) {
  const doc = window.document;
  const docEl = doc.documentElement;
  let requestFullScreen =
    docEl.requestFullscreen ||
    docEl.mozRequestFullScreen ||
    docEl.webkitRequestFullScreen ||
    docEl.msRequestFullscreen;
  let cancelFullScreen =
    doc.exitFullscreen ||
    doc.mozCancelFullScreen ||
    doc.webkitExitFullscreen ||
    doc.msExitFullscreen;
  if (!isFullScreen()) {
    requestFullScreen.call(docEl);
  } else {
    cancelFullScreen.call(doc);
  }
}

function fullScreenChangeListener(event) {
  fullScreenSwitch.checked = isFullScreen();
}

function pageVisibilitySwitchListener(event) {
  let dataPageId =
    event.currentTarget.attributes.getNamedItem("data-page-id").value;
  localStorage.setItem(
    dataPageId + "-visibility",
    Number(event.currentTarget.checked),
  );
  showPage(dataPageId, event.currentTarget.checked);
}

function applyTheme(isDarkMode) {
  const theme = isDarkMode ? "dark" : "light";
  themeLink.setAttribute("href", `static/css/${theme}.css`);
  themeMeta.setAttribute("content", isDarkMode ? "#111111" : "#ffffff");
  darkModeSwitch.checked = isDarkMode;
}

async function sendKey(key) {
  try {
    const response = await fetch("/api", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ key }),
    });
    if (!response.ok) {
      const errorText = await response.text();
      showErrorModal(
        `${response.status}: ${response.statusText}: ${errorText}`,
      );
    }
  } catch (error) {
    console.error("Error sending key:", error);
    showErrorModal("Error sending key:", error.message);
  }
}

function isFullScreen() {
  return (
    !window.document.fullscreenElement &&
    !window.document.mozFullScreenElement &&
    !window.document.webkitFullscreenElement &&
    !window.document.msFullscreenElement
  );
}

function showPage(id, show) {
  document.getElementById(id).classList.toggle("d-none", !show);
  document.getElementById(id + "-tab").classList.toggle("d-none", !show);
  document.getElementById(id + "-page-visibility-switch").checked = show;
}

function showErrorModal(message) {
  const modalId = "modal_" + Math.random().toString(36).substring(2, 9);
  const modalHtml = `
        <div class="modal fade" id="${modalId}" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title text-danger">Error</h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                ${message}
            </div>
            </div>
        </div>
        </div>
    `;
  const wrapper = document.createElement("div");
  wrapper.innerHTML = modalHtml;
  document.body.appendChild(wrapper);
  const modalElement = document.getElementById(modalId);
  const modal = new bootstrap.Modal(modalElement);
  modal.show();
  modalElement.addEventListener("hidden.bs.modal", () => {
    modalElement.remove();
  });
}
