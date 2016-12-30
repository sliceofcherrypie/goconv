package goconv
"""
Função para converter milhas em kilometros
"""
func Ml_to_km(ml float64) float64 {
	km := ml * Ml
	return km
}
"""
Função para kilometros milhas em milhas
"""
func Kl_to_ml(km float64) float64 {
	ml := km / Ml
	return ml
}
