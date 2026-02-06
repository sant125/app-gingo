package data

import "github.com/santzin/gin-tattoo/internal/models"

// Styles holds all in-memory tattoo style records.
var Styles = []models.Style{
	{
		ID:          1,
		Name:        "Old School",
		Description: "Estilo clássico americano com linhas grossas, cores sólidas e motivos icônicos como âncoras, rosas e pássaros. Surgiu no início do século XX com marinheiros.",
		Origin:      "Estados Unidos",
		Popularity:  "high",
	},
	{
		ID:          2,
		Name:        "New School",
		Description: "Evolução do Old School com paleta de cores vibrantes e exageradas, influências de cartoons e cultura pop. Traços caricaturais e composições dinâmicas.",
		Origin:      "Estados Unidos / Europa",
		Popularity:  "medium",
	},
	{
		ID:          3,
		Name:        "Realismo",
		Description: "Reproduz fotografias ou pinturas com altíssimo nível de detalhe e sombreado. Exige grande habilidade técnica do tatuador para criar ilusão tridimensional.",
		Origin:      "Global",
		Popularity:  "high",
	},
	{
		ID:          4,
		Name:        "Blackwork",
		Description: "Utiliza exclusivamente tinta preta, com grandes áreas sólidas, padrões tribais e geométricos. Inspira-se em tradições indígenas de diversas culturas.",
		Origin:      "Polinésia / Global",
		Popularity:  "high",
	},
	{
		ID:          5,
		Name:        "Aquarela",
		Description: "Imita a técnica de pintura em aquarela com splashes de cor, sem contornos definidos. Cria efeito fluido e artístico único na pele.",
		Origin:      "Europa",
		Popularity:  "medium",
	},
	{
		ID:          6,
		Name:        "Japonês",
		Description: "Tradição milenar com motivos como dragões, carpas koi, flores de cerejeira e guerreiros. Usa técnica de degradê chamada bokashi e composições que fluem pelo corpo.",
		Origin:      "Japão",
		Popularity:  "high",
	},
	{
		ID:          7,
		Name:        "Geométrico",
		Description: "Baseado em formas geométricas precisas, simetria e padrões matemáticos. Pode combinar linhas finas com áreas sólidas para criar mandalas e poliedros.",
		Origin:      "Global",
		Popularity:  "high",
	},
	{
		ID:          8,
		Name:        "Fineline",
		Description: "Utiliza agulhas finíssimas para criar traços delicados e detalhes minuciosos. Popular para tatuagens pequenas, lettering e retratos sutis.",
		Origin:      "Estados Unidos",
		Popularity:  "high",
	},
}

// Curiosities holds all in-memory tattoo curiosity records.
var Curiosities = []models.Curiosity{
	{
		ID:       1,
		Title:    "A tinta que persiste",
		Content:  "A tinta da tatuagem não fica na epiderme (camada externa), mas sim na derme — a camada intermediária da pele. Os macrófagos do sistema imune tentam eliminar as partículas, mas como são grandes demais, ficam presos nessa camada, mantendo a tatuagem visível por décadas.",
		Category: "science",
	},
	{
		ID:       2,
		Title:    "Tatuagens na Antiguidade",
		Content:  "O mais antigo registro humano de tatuagem pertence a Ötzi, o Homem de Gelo, que viveu há mais de 5.300 anos. Seu corpo preservado no gelo possuía 61 tatuagens, a maioria em articulações — possivelmente usadas como terapia para dores.",
		Category: "history",
	},
	{
		ID:       3,
		Title:    "Cicatrização em fases",
		Content:  "A cicatrização de uma tatuagem passa por três fases: fase inflamatória (1-3 dias, com vermelhidão e inchaço), fase proliferativa (3-21 dias, com descamação e coceira) e fase de remodelação (até 6 meses, quando as cores se estabilizam definitivamente).",
		Category: "science",
	},
	{
		ID:       4,
		Title:    "Significados culturais",
		Content:  "Para os Maoris da Nova Zelândia, o Ta Moko (tatuagem facial) é uma expressão de identidade, linhagem e status social. Cada padrão é único e conta a história da pessoa. Reproduzir um Ta Moko sem origem maori é considerado desrespeitoso culturalmente.",
		Category: "culture",
	},
	{
		ID:       5,
		Title:    "Composição da tinta",
		Content:  "As tintas modernas para tatuagem são compostas por pigmentos (orgânicos ou inorgânicos), carreadores líquidos (água, álcool, glicerina) e aditivos estabilizadores. As cores claras como amarelo e branco geralmente desbotam mais rápido que preto e azul escuro.",
		Category: "science",
	},
	{
		ID:       6,
		Title:    "Tatuagens no Egito Antigo",
		Content:  "Múmias egípcias femininas datadas de 2000 a.C. apresentam tatuagens em padrões de pontos e linhas, especialmente no abdômen, coxas e seios. Pesquisadores acreditam que tinham função ritual e de fertilidade, possivelmente relacionadas à deusa Hathor.",
		Category: "history",
	},
	{
		ID:       7,
		Title:    "A máquina de tatuar",
		Content:  "A primeira máquina elétrica de tatuagem foi patenteada por Samuel O'Reilly em 1891, baseada em uma invenção de Thomas Edison para gravar superfícies. Antes disso, tatuadores usavam agulhas presas em hastes de madeira ou osso, tingidas à mão.",
		Category: "history",
	},
	{
		ID:       8,
		Title:    "Remoção a laser",
		Content:  "O laser de remoção de tatuagens funciona emitindo pulsos de luz de alta energia que fragmentam as partículas de tinta em pedaços microscópicos. O sistema imunológico então elimina esses fragmentos gradualmente. Cores como verde e amarelo são as mais difíceis de remover.",
		Category: "science",
	},
	{
		ID:       9,
		Title:    "Tatuagens na arte japonesa",
		Content:  "O estilo irezumi japonês tem raízes nos woodblock prints (ukiyo-e) do período Edo (1603-1868). Mestres tatuadores eram frequentemente artistas gráficos que adaptaram seus desenhos para a pele. O dragão, a carpa koi e a fênix são os motivos mais icônicos.",
		Category: "art",
	},
	{
		ID:       10,
		Title:    "Dor e endorfina",
		Content:  "Durante o processo de tatuagem, o corpo libera endorfinas — neurotransmissores que atuam como analgésicos naturais — em resposta à dor das agulhas. Isso pode criar uma sensação de euforia, que alguns tatuados descrevem como 'vício' nas sessões.",
		Category: "science",
	},
}
