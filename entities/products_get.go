package bling

type Categoria struct {
	ID int `json:"id,omitempty"`
}

type FornecedorContato struct {
	ID   int    `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
}

type Fornecedor struct {
	ID          int               `json:"id,omitempty"`
	Contato     FornecedorContato `json:"contato,omitempty"`
	Codigo      string            `json:"codigo,omitempty"`
	PrecoCusto  float64           `json:"precoCusto,omitempty"`
	PrecoCompra float64           `json:"precoCompra,omitempty"`
}

type Estoque struct {
	Minimo            int    `json:"minimo,omitempty"`
	Maximo            int    `json:"maximo,omitempty"`
	Crossdocking      int    `json:"crossdocking,omitempty"`
	Localizacao       string `json:"localizacao,omitempty"`
	SaldoVirtualTotal int    `json:"saldoVirtualTotal"`
}

type Dimensoes struct {
	Largura       float64 `json:"largura,omitempty"`
	Altura        float64 `json:"altura,omitempty"`
	Profundidade  float64 `json:"profundidade,omitempty"`
	UnidadeMedida int     `json:"unidadeMedida,omitempty"`
}

type Tributacao struct {
	Origem              int     `json:"origem,omitempty"`
	NFCI                string  `json:"nFCI,omitempty"`
	NCM                 string  `json:"ncm,omitempty"`
	CEST                string  `json:"cest,omitempty"`
	CodigoListaServicos string  `json:"codigoListaServicos,omitempty"`
	SpedTipoItem        string  `json:"spedTipoItem,omitempty"`
	CodigoItem          string  `json:"codigoItem,omitempty"`
	PercentualTributos  float64 `json:"percentualTributos,omitempty"`
}

type ImagemInterna struct {
	LinkMiniatura string `json:"linkMiniatura,omitempty"`
	Validade      string `json:"validade,omitempty"`
	Ordem         int    `json:"ordem,omitempty"`
}

type Midia struct {
	Video   string          `json:"url,omitempty"`
	Imagens []ImagemInterna `json:"internas,omitempty"`
}

type Produto struct {
	ID                         int        `json:"id"`
	Nome                       string     `json:"nome"`
	Codigo                     string     `json:"codigo"`
	Preco                      float64    `json:"preco"`
	Tipo                       string     `json:"tipo"`
	Situacao                   string     `json:"situacao"`
	Formato                    string     `json:"formato"`
	DescricaoCurta             string     `json:"descricaoCurta,omitempty"`
	ImagemURL                  string     `json:"imagemURL,omitempty"`
	DataValidade               string     `json:"dataValidade,omitempty"`
	Unidade                    string     `json:"unidade,omitempty"`
	PesoLiquido                float64    `json:"pesoLiquido,omitempty"`
	PesoBruto                  float64    `json:"pesoBruto,omitempty"`
	Volumes                    int        `json:"volumes,omitempty"`
	ItensPorCaixa              int        `json:"itensPorCaixa,omitempty"`
	GTIN                       string     `json:"gtin,omitempty"`
	GTINEmbalagem              string     `json:"gtinEmbalagem,omitempty"`
	TipoProducao               string     `json:"tipoProducao,omitempty"`
	Condicao                   int        `json:"condicao,omitempty"`
	FreteGratis                bool       `json:"freteGratis,omitempty"`
	Marca                      string     `json:"marca,omitempty"`
	DescricaoComplementar      string     `json:"descricaoComplementar,omitempty"`
	LinkExterno                string     `json:"linkExterno,omitempty"`
	Observacoes                string     `json:"observacoes,omitempty"`
	DescricaoEmbalagemDiscreta string     `json:"descricaoEmbalagemDiscreta,omitempty"`
	Categoria                  Categoria  `json:"categoria,omitempty"`
	Estoque                    Estoque    `json:"estoque"`
	Fornecedor                 Fornecedor `json:"fornecedor,omitempty"`
	ActionEstoque              string     `json:"actionEstoque,omitempty"`
	Dimensoes                  Dimensoes  `json:"dimensoes,omitempty"`
	Tributacao                 Tributacao `json:"tributacao,omitempty"`
	Midia                      Midia      `json:"midia,omitempty"`
}
