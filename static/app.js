/* ── Gin Tattoo — Frontend ────────────────────────────────────────────────── */

const BASE = '/api/v1';

// ── Utilities ─────────────────────────────────────────────────────────────────

async function fetchJSON(url) {
  const res = await fetch(url);
  if (!res.ok) throw new Error(`HTTP ${res.status}`);
  return res.json();
}

function openModal(html) {
  document.getElementById('modal-body').innerHTML = html;
  document.getElementById('modal').classList.remove('hidden');
}

function closeModal() {
  document.getElementById('modal').classList.add('hidden');
}

document.getElementById('modal-close').addEventListener('click', closeModal);
document.getElementById('modal').addEventListener('click', (e) => {
  if (e.target === document.getElementById('modal')) closeModal();
});
document.addEventListener('keydown', (e) => { if (e.key === 'Escape') closeModal(); });

// ── Styles ────────────────────────────────────────────────────────────────────

async function loadStyles() {
  const grid = document.getElementById('styles-grid');
  try {
    const styles = await fetchJSON(`${BASE}/styles`);
    grid.innerHTML = '';
    styles.forEach((s) => {
      const card = document.createElement('div');
      card.className = 'card';
      card.innerHTML = `
        <div class="card-name">${s.name}</div>
        <div class="card-origin">${s.origin}</div>
        <span class="badge badge-${s.popularity}">${s.popularity}</span>
      `;
      card.addEventListener('click', () => showStyle(s.id));
      grid.appendChild(card);
    });
  } catch (err) {
    grid.innerHTML = `<p class="loading">Erro ao carregar estilos: ${err.message}</p>`;
  }
}

async function showStyle(id) {
  try {
    const s = await fetchJSON(`${BASE}/styles/${id}`);
    openModal(`
      <h3>${s.name}</h3>
      <p class="meta">Origem: ${s.origin} &nbsp;|&nbsp; Popularidade: ${s.popularity}</p>
      <p>${s.description}</p>
    `);
  } catch (err) {
    openModal(`<p>Erro: ${err.message}</p>`);
  }
}

// ── Curiosities ───────────────────────────────────────────────────────────────

async function loadCuriosities() {
  const list = document.getElementById('curiosities-list');
  try {
    const curiosities = await fetchJSON(`${BASE}/curiosities`);
    list.innerHTML = '';
    curiosities.forEach((c) => {
      const item = document.createElement('div');
      item.className = 'curiosity-item';
      item.innerHTML = `
        <div class="curiosity-title">${c.title}</div>
        <div class="curiosity-category">${c.category}</div>
        <div class="curiosity-preview">${c.content}</div>
      `;
      item.addEventListener('click', () => showCuriosity(c.id));
      list.appendChild(item);
    });
  } catch (err) {
    list.innerHTML = `<p class="loading">Erro ao carregar curiosidades: ${err.message}</p>`;
  }
}

async function showCuriosity(id) {
  try {
    const c = await fetchJSON(`${BASE}/curiosities/${id}`);
    openModal(`
      <h3>${c.title}</h3>
      <p class="meta">Categoria: ${c.category}</p>
      <p>${c.content}</p>
    `);
  } catch (err) {
    openModal(`<p>Erro: ${err.message}</p>`);
  }
}

// ── Init ──────────────────────────────────────────────────────────────────────
loadStyles();
loadCuriosities();
