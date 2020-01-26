TEXMFHOME=$(shell kpsewhich -var-value TEXMFHOME)

all: e1.pdf

e1.pdf: e1.tex
	latexrun $<

.PHONY: all deps clean

deps:
	git clone https://github.com/rpgtex/DND-5e-LaTeX-Template.git "${TEXMFHOME}/tex/latex/dnd"

clean:
	$(RM) -r latex.out
	$(RM) e1.pdf
