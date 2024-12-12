package service


// TODO: add vaildation for urls ? idk

// AddUrl implements Service.
func (s *SrvImpl) AddUrl(url string) (string, error) {
  // give this request an id.
  // add this url to a queue in SrvImpl.
  //
	panic("unimplemented")
}

// GetFullURL implements Service.
func (s *SrvImpl) GetFullURL(key string) (string, error) {
  // check if key exists in cache
  if v, ok := s.cache.Get(key); ok {
    return v, nil
  }
  
  
  

	panic("unimplemented")
}
